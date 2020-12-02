package mp

import (
	"context"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
	"github.com/tidwall/gjson"
)

// MP 微信小程序
type MP struct {
	appid          string
	appsecret      string
	token          string
	encodingAESKey string
	nonce          func(size int) string
	client         wx.Client
}

// New returns new wechat mini program
func New(appid, appsecret string) *MP {
	return &MP{
		appid:     appid,
		appsecret: appsecret,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client: wx.NewHTTPClient(),
	}
}

// SetServerConfig 设置服务器配置
func (mp *MP) SetServerConfig(token, encodingAESKey string) {
	mp.token = token
	mp.encodingAESKey = encodingAESKey
}

// Code2Session 获取小程序授权的session_key
func (mp *MP) Code2Session(ctx context.Context, code string, options ...wx.HTTPOption) (*AuthSession, error) {
	resp, err := mp.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", Code2SessionURL, mp.appid, mp.appsecret, code), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	session := new(AuthSession)

	if err = json.Unmarshal(resp, session); err != nil {
		return nil, err
	}

	return session, nil
}

// AccessToken 获取小程序的access_token
func (mp *MP) AccessToken(ctx context.Context, options ...wx.HTTPOption) (*AccessToken, error) {
	resp, err := mp.client.Get(ctx, fmt.Sprintf("%s?appid=%s&secret=%s&grant_type=client_credential", AccessTokenURL, mp.appid, mp.appsecret), options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	token := new(AccessToken)

	if err = json.Unmarshal(resp, token); err != nil {
		return nil, err
	}

	return token, nil
}

// DecryptAuthInfo 解密授权信息
func (mp *MP) DecryptAuthInfo(sessionKey, iv, encryptedData string, dest AuthInfo) error {
	key, err := base64.StdEncoding.DecodeString(sessionKey)

	if err != nil {
		return err
	}

	ivb, err := base64.StdEncoding.DecodeString(iv)

	if err != nil {
		return err
	}

	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)

	if err != nil {
		return err
	}

	cbc := wx.NewAESCBCCrypto(key, ivb)

	b, err := cbc.Decrypt(cipherText, wx.PKCS7)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, dest); err != nil {
		return err
	}

	if dest.AppID() != mp.appid {
		return fmt.Errorf("appid mismatch, want: %s, got: %s", mp.appid, dest.AppID())
	}

	return nil
}

// Do exec action
func (mp *MP) Do(ctx context.Context, accessToken string, action wx.Action, options ...wx.HTTPOption) error {
	var (
		resp []byte
		err  error
	)

	switch action.Method() {
	case wx.MethodGet:
		resp, err = mp.client.Get(ctx, action.URL()(accessToken), options...)
	case wx.MethodPost:
		resp, err = mp.client.Post(ctx, action.URL()(accessToken), action.Body(), options...)
	case wx.MethodUpload:
		resp, err = mp.client.Upload(ctx, action.URL()(accessToken), action.Body(), options...)
	}

	if err != nil {
		return err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return errors.New(r.Get("errmsg").String())
	}

	if action.Decode() == nil {
		return nil
	}

	return action.Decode()(resp)
}

// DecryptEventMessage 事件消息解密
func (mp *MP) DecryptEventMessage(msgEncrypt string) (*event.Message, error) {
	b, err := event.Decrypt(mp.appid, mp.encodingAESKey, msgEncrypt)

	if err != nil {
		return nil, err
	}

	msg := new(event.Message)

	if err = xml.Unmarshal(b, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

// VerifyServer 验证消息来自微信服务器（若验证成功，请原样返回echostr参数内容）
func (mp *MP) VerifyServer(signature, timestamp, nonce string) bool {
	signItems := []string{mp.token, timestamp, nonce}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))
	signStr := hex.EncodeToString(h.Sum(nil))

	return signStr == signature
}

// VerifyEvent 验证事件签名（注意：使用 msg_signature）
func (mp *MP) VerifyEvent(msgSignature, timestamp, nonce, msgEncrypt string) bool {
	signItems := []string{mp.token, timestamp, nonce, msgEncrypt}

	sort.Strings(signItems)

	h := sha1.New()
	h.Write([]byte(strings.Join(signItems, "")))
	signStr := hex.EncodeToString(h.Sum(nil))

	return signStr == msgSignature
}