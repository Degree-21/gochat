package minip

import (
	"encoding/json"

	"github.com/shenghui0779/gochat/urls"
	"github.com/shenghui0779/gochat/wx"
)

// Gender 性别
type Gender int

const (
	GenderUnknown Gender = 0 // 未知
	GenderMale    Gender = 1 // 男性
	GenderFemale  Gender = 2 // 女性
)

// AuthSession 小程序授权Session
type AuthSession struct {
	SessionKey string `json:"session_key"`
	OpenID     string `json:"openid"`
	UnionID    string `json:"unionid"`
}

// AccessToken 小程序access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthInfo 小程序授权信息
type AuthInfo interface {
	AppID() string
}

// UserInfo 用户信息
type UserInfo struct {
	OpenID    string    `json:"openId"`
	Language  string    `json:"language"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	AvatarURL string    `json:"avatarUrl"`
	NickName  string    `json:"nickName"`
	Gender    Gender    `json:"gender"`
	Country   string    `json:"country"`
	UnionID   string    `json:"unionId"`
	WaterMark WaterMark `json:"watermark"`
}

func (u *UserInfo) AppID() string {
	return u.WaterMark.AppID
}

// PhoneInfo 用户手机号绑定信息
type PhoneInfo struct {
	PhoneNumber     string    `json:"phoneNumber"`
	PurePhoneNumber string    `json:"purePhoneNumber"`
	CountryCode     string    `json:"countryCode"`
	WaterMark       WaterMark `json:"watermark"`
}

func (p *PhoneInfo) AppID() string {
	return p.WaterMark.AppID
}

// WaterMark 水印
type WaterMark struct {
	Timestamp int64  `json:"timestamp"`
	AppID     string `json:"appid"`
}

// ResultPaidUnionID 支付用户unionid
type ResultPaidUnionID struct {
	UnionID string `json:"unionid"`
}

// GetPaidUnionIDByTransactionID 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByTransactionID(openid, transactionID string, result *ResultPaidUnionID) wx.Action {
	return wx.NewGetAction(urls.MinipPaidUnion,
		wx.WithQuery("openid", openid),
		wx.WithQuery("transaction_id", transactionID),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}

// GetPaidUnionIDByOutTradeNO 用户支付完成后，获取该用户的 UnionId，无需用户授权
func GetPaidUnionIDByOutTradeNO(openid, mchid, outTradeNO string, result *ResultPaidUnionID) wx.Action {
	return wx.NewGetAction(urls.MinipPaidUnion,
		wx.WithQuery("openid", openid),
		wx.WithQuery("mch_id", mchid),
		wx.WithQuery("out_trade_no", outTradeNO),
		wx.WithDecode(func(resp []byte) error {
			return json.Unmarshal(resp, result)
		}),
	)
}
