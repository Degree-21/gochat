package mchv3

import (
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/shenghui0779/gochat/wx"
)

type MchV3 struct {
	appid  string
	mchid  string
	apikey string
	mcert  wx.RSACert
	wcert  map[string]wx.RSACert
	nonce  func(size int) string
	client wx.HTTPClient
}

// New returns new wechat pay
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

// LoadMCertFromPemFile 加载商户证书
func (mch *MchV3) LoadMCertFromPemFile(certFile, keyFile string) error {
	cert, err := wx.NewRSACertFromPemFile(certFile, keyFile)

	if err != nil {
		return err
	}

	mch.mcert = cert

	return nil
}

// LoadMCertFromPemBlock 加载商户证书
func (mch *MchV3) LoadMCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := wx.NewRSACertFromPemBlock(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.mcert = cert

	return nil
}

// LoadWCertFromPemFile 加载平台证书
func (mch *MchV3) LoadWCertFromPemFile(certFile, keyFile string) error {
	cert, err := wx.NewRSACertFromPemFile(certFile, keyFile)

	if err != nil {
		return err
	}

	mch.wcert[cert.SerialNumber()] = cert

	return nil
}

// LoadWCertFromPemBlock 加载平台证书
func (mch *MchV3) LoadWCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := wx.NewRSACertFromPemBlock(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.wcert[cert.SerialNumber()] = cert

	return nil
}
