package mchv3

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"
	"sync"

	"github.com/shenghui0779/gochat/wx"
)

// MchV3 微信支付V3
type MchV3 struct {
	appid  string
	mchid  string
	apikey string
	mcert  wx.RSACert
	wcert  sync.Map
	prvkey wx.RSAKey
	nonce  func(size int) string
	client wx.HTTPClient
	mutex  sync.RWMutex
}

// New returns new wechat pay v3
func New(appid, mchid, apikey string) *MchV3 {
	return &MchV3{
		appid:  appid,
		mchid:  mchid,
		apikey: apikey,
		nonce: func(size int) string {
			nonce := make([]byte, size/2)
			io.ReadFull(rand.Reader, nonce)

			return hex.EncodeToString(nonce)
		},
		client: wx.NewHTTPClient(),
	}
}

// LoadCertificate 加载证书
func (mch *MchV3) LoadCertificate(options ...CertOption) error {
	for _, f := range options {
		if err := f(mch); err != nil {
			return err
		}
	}

	return nil
}

// AppID returns appid
func (mch *MchV3) AppID() string {
	return mch.appid
}

// MchID returns mchid
func (mch *MchV3) MchID() string {
	return mch.mchid
}

// ApiKey returns apikey
func (mch *MchV3) ApiKey() string {
	return mch.apikey
}

// CertOption 证书选项
type CertOption func(mch *MchV3) error

// WithMchCertBlock 通过商户证书文本内容加载证书
func WithMchCertBlock(certPEMBlock []byte) CertOption {
	return func(mch *MchV3) error {
		cert, err := wx.NewRSACert(certPEMBlock)

		if err != nil {
			return err
		}

		mch.mcert = cert

		return nil
	}
}

// WithMchCertFile 通过商户证书文件加载证书
func WithMchCertFile(certPEMFile string) CertOption {
	return func(mch *MchV3) error {
		certPEMBlock, err := os.ReadFile(certPEMFile)

		if err != nil {
			return err
		}

		cert, err := wx.NewRSACert(certPEMBlock)

		if err != nil {
			return err
		}

		mch.mcert = cert

		return nil
	}
}

// WithWechatCertBlock 通过平台证书文本内容加载证书
func WithWechatCertBlock(certPEMBlock []byte) CertOption {
	return func(mch *MchV3) error {
		cert, err := wx.NewRSACert(certPEMBlock)

		if err != nil {
			return err
		}

		mch.wcert.Store(cert.SerialNumber(), cert)

		return nil
	}
}

// WithWechatCertFile 通过平台证书文件加载证书
func WithWechatCertFile(certPEMFile string) CertOption {
	return func(mch *MchV3) error {
		certPEMBlock, err := os.ReadFile(certPEMFile)

		if err != nil {
			return err
		}

		cert, err := wx.NewRSACert(certPEMBlock)

		if err != nil {
			return err
		}

		mch.wcert.Store(cert.SerialNumber(), cert)

		return nil
	}
}

// WithPrivateKeyBlock 通过商户API私钥文本内容加载私钥
func WithPrivateKeyBlock(keyPEMBlock []byte) CertOption {
	return func(mch *MchV3) error {
		key, err := wx.NewRSAKey(keyPEMBlock)

		if err != nil {
			return err
		}

		mch.prvkey = key

		return nil
	}
}

// WithPrivateKeyFile 通过商户API私钥文件加载私钥
func WithPrivateKeyFile(keyPEMFile string) CertOption {
	return func(mch *MchV3) error {
		keyPEMBlock, err := os.ReadFile(keyPEMFile)

		if err != nil {
			return err
		}

		key, err := wx.NewRSAKey(keyPEMBlock)

		if err != nil {
			return err
		}

		mch.prvkey = key

		return nil
	}
}
