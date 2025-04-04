package mch

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/shenghui0779/yiigo"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Mch 微信支付
type Mch struct {
	appid  string
	mchid  string
	apikey string
	nonce  func() string
	client wx.Client
	tlscli wx.Client
}

// New returns new wechat pay
// [证书参考](https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=4_3)
func New(appid, mchid, apikey string, certs ...tls.Certificate) *Mch {
	return &Mch{
		appid:  appid,
		mchid:  mchid,
		apikey: apikey,
		nonce: func() string {
			return wx.Nonce(16)
		},
		client: wx.DefaultClient(),
		tlscli: wx.DefaultClient(certs...),
	}
}

// SetClient sets options for wechat client
func (mch *Mch) SetClient(options ...wx.ClientOption) {
	mch.client.Set(options...)
}

// SetTLSClient sets options for wechat tls client
func (mch *Mch) SetTLSClient(options ...wx.ClientOption) {
	mch.tlscli.Set(options...)
}

// AppID returns appid
func (mch *Mch) AppID() string {
	return mch.appid
}

// MchID returns mchid
func (mch *Mch) MchID() string {
	return mch.mchid
}

// ApiKey returns apikey
func (mch *Mch) ApiKey() string {
	return mch.apikey
}

// Do exec action
func (mch *Mch) Do(ctx context.Context, action wx.Action, options ...yiigo.HTTPOption) (wx.WXML, error) {
	m, err := action.WXML(mch.appid, mch.mchid, mch.nonce())

	if err != nil {
		return nil, err
	}

	// 签名
	if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
		m["sign"] = mch.SignWithHMacSHA256(m, true)
	} else {
		m["sign"] = mch.SignWithMD5(m, true)
	}

	if len(action.Method()) == 0 {
		if len(action.URL()) == 0 {
			return m, nil
		}

		query := url.Values{}

		for k, v := range m {
			query.Add(k, v)
		}

		return wx.WXML{"entrust_url": fmt.Sprintf("%s?%s", action.URL(), query.Encode())}, nil
	}

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	var resp []byte

	if action.IsTLS() {
		resp, err = mch.tlscli.Do(ctx, action.Method(), action.URL(), body, options...)
	} else {
		resp, err = mch.client.Do(ctx, action.Method(), action.URL(), body, options...)
	}

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	// 签名验证
	if err := mch.VerifyWXMLResult(result); err != nil {
		return nil, err
	}

	return result, nil
}

// APPAPI 用于APP拉起支付
func (mch *Mch) APPAPI(prepayID string) wx.WXML {
	m := wx.WXML{
		"appid":     mch.appid,
		"partnerid": mch.mchid,
		"prepayid":  prepayID,
		"package":   "Sign=WXPay",
		"noncestr":  mch.nonce(),
		"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["sign"] = mch.SignWithMD5(m, true)

	return m
}

// JSAPI 用于JS拉起支付
func (mch *Mch) JSAPI(prepayID string) wx.WXML {
	m := wx.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  SignMD5,
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = mch.SignWithMD5(m, true)

	return m
}

// MinipRedpackJSAPI 小程序领取红包
func (mch *Mch) MinipRedpackJSAPI(pkg string) wx.WXML {
	m := wx.WXML{
		"appId":     mch.appid,
		"nonceStr":  mch.nonce(),
		"package":   url.QueryEscape(pkg),
		"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
	}

	m["paySign"] = mch.SignWithMD5(m, false)

	delete(m, "appId")
	m["signType"] = SignMD5

	return m
}

// DownloadBill 下载交易账单
// 账单日期格式：20140603
func (mch *Mch) DownloadBill(ctx context.Context, billDate, billType string) ([]byte, error) {
	m := wx.WXML{
		"appid":     mch.appid,
		"mch_id":    mch.mchid,
		"bill_date": billDate,
		"bill_type": billType,
		"nonce_str": mch.nonce(),
	}

	m["sign"] = mch.SignWithMD5(m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.client.Do(ctx, http.MethodPost, urls.MchDownloadBill, body, yiigo.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// DownloadFundFlow 下载资金账单
// 账单日期格式：20140603
func (mch *Mch) DownloadFundFlow(ctx context.Context, billDate, accountType string) ([]byte, error) {
	m := wx.WXML{
		"appid":        mch.appid,
		"mch_id":       mch.mchid,
		"bill_date":    billDate,
		"account_type": accountType,
		"nonce_str":    mch.nonce(),
	}

	m["sign"] = mch.SignWithHMacSHA256(m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.tlscli.Do(ctx, http.MethodPost, urls.MchDownloadFundFlow, body, yiigo.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// BatchQueryComment 拉取订单评价数据
// 时间格式：yyyyMMddHHmmss
// 默认一次且最多拉取200条
func (mch *Mch) BatchQueryComment(ctx context.Context, beginTime, endTime string, offset int, limit ...int) ([]byte, error) {
	m := wx.WXML{
		"appid":      mch.appid,
		"mch_id":     mch.mchid,
		"begin_time": beginTime,
		"end_time":   endTime,
		"offset":     strconv.Itoa(offset),
		"nonce_str":  mch.nonce(),
	}

	if len(limit) != 0 {
		m["limit"] = strconv.Itoa(limit[0])
	}

	m["sign"] = mch.SignWithHMacSHA256(m, true)

	body, err := wx.FormatMap2XML(m)

	if err != nil {
		return nil, err
	}

	resp, err := mch.tlscli.Do(ctx, http.MethodPost, urls.MchBatchQueryComment, body, yiigo.WithHTTPClose())

	if err != nil {
		return nil, err
	}

	// XML解析
	result, err := wx.ParseXML2Map(resp)

	if err != nil {
		return nil, err
	}

	if len(result) != 0 && result["return_code"] != ResultSuccess {
		return nil, errors.New(result["return_msg"])
	}

	return resp, nil
}

// SignWithMD5 生成MD5签名
func (mch *Mch) SignWithMD5(m wx.WXML, toUpper bool) string {
	h := md5.New()
	h.Write([]byte(mch.buildSignStr(m)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// SignWithHMacSHA256 生成HMAC-SHA256签名
func (mch *Mch) SignWithHMacSHA256(m wx.WXML, toUpper bool) string {
	h := hmac.New(sha256.New, []byte(mch.apikey))
	h.Write([]byte(mch.buildSignStr(m)))

	sign := hex.EncodeToString(h.Sum(nil))

	if toUpper {
		sign = strings.ToUpper(sign)
	}

	return sign
}

// VerifyWXMLResult 微信请求/回调通知签名验证
func (mch *Mch) VerifyWXMLResult(m wx.WXML) error {
	if wxsign, ok := m["sign"]; ok {
		signature := ""

		if v, ok := m["sign_type"]; ok && v == SignHMacSHA256 {
			signature = mch.SignWithHMacSHA256(m, true)
		} else {
			signature = mch.SignWithMD5(m, true)
		}

		if wxsign != signature {
			return fmt.Errorf("signature verified failed, want: %s, got: %s", signature, wxsign)
		}
	}

	if appid, ok := m["appid"]; ok {
		if appid != mch.appid {
			return fmt.Errorf("appid mismatch, want: %s, got: %s", mch.appid, m["appid"])
		}
	}

	if mchid, ok := m["mch_id"]; ok {
		if mchid != mch.mchid {
			return fmt.Errorf("mchid mismatch, want: %s, got: %s", mch.mchid, m["mch_id"])
		}
	}

	return nil
}

// DecryptWithAES256ECB AES-256-ECB解密（主要用于退款结果通知）
func (mch *Mch) DecryptWithAES256ECB(encrypt string) (wx.WXML, error) {
	cipherText, err := base64.StdEncoding.DecodeString(encrypt)

	if err != nil {
		return nil, err
	}

	h := md5.New()
	h.Write([]byte(mch.apikey))

	ecb := yiigo.NewECBCrypto([]byte(hex.EncodeToString(h.Sum(nil))), yiigo.PKCS7)

	plainText, err := ecb.Decrypt(cipherText)

	if err != nil {
		return nil, err
	}

	return wx.ParseXML2Map(plainText)
}

// Sign 生成签名
func (mch *Mch) buildSignStr(m wx.WXML) string {
	l := len(m)

	ks := make([]string, 0, l)
	kvs := make([]string, 0, l)

	for k := range m {
		if k == "sign" {
			continue
		}

		ks = append(ks, k)
	}

	sort.Strings(ks)

	for _, k := range ks {
		if v, ok := m[k]; ok && v != "" {
			kvs = append(kvs, fmt.Sprintf("%s=%s", k, v))
		}
	}

	kvs = append(kvs, fmt.Sprintf("key=%s", mch.apikey))

	return strings.Join(kvs, "&")
}
