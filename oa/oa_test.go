package oa

import (
	"context"
	"encoding/xml"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/event"
	"github.com/shenghui0779/gochat/wx"
	"github.com/stretchr/testify/assert"
)

func TestAuthURL(t *testing.T) {
	oa := New("APPID", "APPSECRET")
	oa.nonce = func(size int) string {
		return "STATE"
	}

	assert.Equal(t, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=RedirectURL&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect", oa.AuthURL(ScopeSnsapiBase, "RedirectURL"))
	assert.Equal(t, "https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=RedirectURL&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect", oa.AuthURL(ScopeSnsapiUser, "RedirectURL"))
}

func TestCode2AuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=APPSECRET&code=CODE&grant_type=authorization_code").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	authToken, err := oa.Code2AuthToken(context.TODO(), "CODE")

	assert.Nil(t, err)
	assert.Equal(t, &AuthToken{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
		ExpiresIn:    7200,
		OpenID:       "OPENID",
		Scope:        "SCOPE",
	}, authToken)
}

func TestRefreshAuthToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN").Return([]byte(`{
		"access_token": "ACCESS_TOKEN",
		"expires_in": 7200,
		"refresh_token": "REFRESH_TOKEN",
		"openid": "OPENID",
		"scope": "SCOPE"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	authToken, err := oa.RefreshAuthToken(context.TODO(), "REFRESH_TOKEN")

	assert.Nil(t, err)
	assert.Equal(t, &AuthToken{
		AccessToken:  "ACCESS_TOKEN",
		RefreshToken: "REFRESH_TOKEN",
		ExpiresIn:    7200,
		OpenID:       "OPENID",
		Scope:        "SCOPE",
	}, authToken)
}

func TestAccessToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET").Return([]byte(`{
		"access_token": "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		"expires_in": 7200
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	accessToken, err := oa.AccessToken(context.TODO())

	assert.Nil(t, err)
	assert.Equal(t, &AccessToken{
		Token:     "39_VzXkFDAJsEVTWbXUZDU3NqHtP6mzcAA7RJvcy1o9e-7fdJ-UuxPYLdBFMiGhpdoeKqVWMGqBe8ldUrMasRv1z_T8RmHKDiybC29wZ_vexHlyQ5YDGb33rff1mBNpOLM9f5nv7oag8UYBSc79ASMcAAADVP",
		ExpiresIn: 7200,
	}, accessToken)
}

func TestDecryptEventMessage(t *testing.T) {
	oa := New("wx1def0e9e5891b338", "APPSECRET")
	oa.SetServerConfig("2faf43d6343a802b6073aae5b3f2f109", "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U")

	msg, err := oa.DecryptEventMessage("GmSmP2C7QlatlbnrXhJHweW5JsW2F1Fr/xmoMBIJNGnZcN/1PoOySJOJNYEC9ttFhaqDrkznaMkDs7s9u7/eOpvqqRn144EBkLdBLxcNbjLRoF4lD3zBGqjPUS9k/U0x/lET35SkYi+ZwRvuSJSzVEfaRmixYep+JmzIYf5k2qT8113wg2tI68+3gUaKZQqq5W/jC7tbWjWX67XgzMW2JdQOs9VnTjJJO292PWkNZxbhzudrvj2Up8NdJbmaDw93Jz/Kcf7qRfdh5h0GFtOoVh7M4bVwTJf94iZU4ZDx1r8/xDxDINRWGJou4Er72cDBCVBK1TUrtwdmb8eWNJ1gSvw53LckULci98+peaSnTFYuaNhgRQqpVQ+CqVjT0+ASRdyMmDomRyUmhBqSsdrGae9pRfP+Dq4tiRoub87T0gGkFTxAXbUZ0ZPxme67ddreWKFCN/V5ypCynDbjkgpIgfPAFpk017ShXc30RRq4qPvPvN/6XUi1HVXSJq8AkgSQ")

	assert.Nil(t, err)
	assert.Equal(t, &event.Message{
		XMLName: xml.Name{
			Space: "",
			Local: "xml",
		},
		ToUserName:   "gh_3ad31c0ba9b5",
		FromUserName: "oB4tA6ANthOfuQ5XSlkdPsWOVUsY",
		CreateTime:   1606902602,
		MsgType:      event.MessageText,
		MsgID:        10086,
		Content:      "ILoveGochat",
	}, msg)
}

func TestVerifyServer(t *testing.T) {
	oa := New("APPID", "APPSECRET")
	oa.SetServerConfig("2faf43d6343a802b6073aae5b3f2f109", "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U")

	assert.True(t, oa.VerifyServer("ffb882ae55647757d3b807ff0e9b6098dfc2bc57", "1606902086", "1246833592"))
}

func TestVerifyEvent(t *testing.T) {
	oa := New("APPID", "APPSECRET")
	oa.SetServerConfig("2faf43d6343a802b6073aae5b3f2f109", "jxAko083VoJ3lcPXJWzcGJ0M1tFVLgdD6qAq57GJY1U")

	assert.True(t, oa.VerifyEvent("4ae269e50a3a4119572be5abe277de261b689f90", "1606903341", "1513542206", "GmSmP2C7QlatlbnrXhJHweW5JsW2F1Fr/xmoMBIJNGnZcN/1PoOySJOJNYEC9ttFhaqDrkznaMkDs7s9u7/eOpvqqRn144EBkLdBLxcNbjLRoF4lD3zBGqjPUS9k/U0x/lET35SkYi+ZwRvuSJSzVEfaRmixYep+JmzIYf5k2qT8113wg2tI68+3gUaKZQqq5W/jC7tbWjWX67XgzMW2JdQOs9VnTjJJO292PWkNZxbhzudrvj2Up8NdJbmaDw93Jz/Kcf7qRfdh5h0GFtOoVh7M4bVwTJf94iZU4ZDx1r8/xDxDINRWGJou4Er72cDBCVBK1TUrtwdmb8eWNJ1gSvw53LckULci98+peaSnTFYuaNhgRQqpVQ+CqVjT0+ASRdyMmDomRyUmhBqSsdrGae9pRfP+Dq4tiRoub87T0gGkFTxAXbUZ0ZPxme67ddreWKFCN/V5ypCynDbjkgpIgfPAFpk017ShXc30RRq4qPvPvN/6XUi1HVXSJq8AkgSQ"))
}

var postBody wx.Body

func TestMain(m *testing.M) {
	postBody = wx.NewPostBody(func() ([]byte, error) {
		return nil, nil
	})

	m.Run()
}