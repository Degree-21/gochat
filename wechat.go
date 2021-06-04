package gochat

import (
	"github.com/shenghui0779/gochat/mch"
	"github.com/shenghui0779/gochat/mchv3"
	"github.com/shenghui0779/gochat/mp"
	"github.com/shenghui0779/gochat/oa"
)

// NewMch 微信商户
func NewMch(appid, mchid, apikey string) *mch.Mch {
	return mch.New(appid, mchid, apikey)
}

// NewMchV3 微信商户V3
func NewMchV3(appid, mchid, apikey string) *mchv3.MchV3 {
	return mchv3.New(appid, mchid, apikey)
}

// NewPub 微信公众号
func NewOA(appid, appsecret string) *oa.OA {
	return oa.New(appid, appsecret)
}

// NewMP 微信小程序
func NewMP(appid, appsecret string) *mp.MP {
	return mp.New(appid, appsecret)
}
