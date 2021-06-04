package mchv3

import "github.com/shenghui0779/gochat/wx"

type MchV3 struct {
	appid     string
	mchid     string
	apikey    string
	nonce     func(size int) string
	client    wx.HTTPClient
	tlsClient wx.HTTPClient
}
