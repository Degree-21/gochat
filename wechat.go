package gochat

import (
	"crypto/tls"

	"github.com/Degree-21/gochat/corp"
	"github.com/Degree-21/gochat/mch"
	"github.com/Degree-21/gochat/minip"
	"github.com/Degree-21/gochat/offia"
)

// NewMch 微信商户
func NewMch(mchid, apikey string, certs ...tls.Certificate) *mch.Mch {
	return mch.New(mchid, apikey, certs...)
}

// NewOffia 微信公众号
func NewOffia(appid, appsecret string) *offia.Offia {
	return offia.New(appid, appsecret)
}

// NewMinip 微信小程序
func NewMinip(appid, appsecret string) *minip.Minip {
	return minip.New(appid, appsecret)
}

// NewCorp 企业微信
func NewCorp(corpid string) *corp.Corp {
	return corp.New(corpid)
}
