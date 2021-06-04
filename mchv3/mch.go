package mchv3

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"os"

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
	certPEMBlock, err := os.ReadFile(certFile)

	if err != nil {
		return err
	}

	keyPEMBlock, err := os.ReadFile(keyFile)

	if err != nil {
		return err
	}

	cert, err := wx.NewRSACert(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.mcert = cert

	return nil
}

// LoadMCertFromPemBlock 加载商户证书
func (mch *MchV3) LoadMCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := wx.NewRSACert(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.mcert = cert

	return nil
}

// LoadWCertFromPemFile 加载平台证书
func (mch *MchV3) LoadWCertFromPemFile(certFile, keyFile string) error {
	certPEMBlock, err := os.ReadFile(certFile)

	if err != nil {
		return err
	}

	keyPEMBlock, err := os.ReadFile(keyFile)

	if err != nil {
		return err
	}

	cert, err := wx.NewRSACert(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.wcert[cert.SerialNumber()] = cert

	return nil
}

// LoadWCertFromPemBlock 加载平台证书
func (mch *MchV3) LoadWCertFromPemBlock(certPEMBlock, keyPEMBlock []byte) error {
	cert, err := wx.NewRSACert(certPEMBlock, keyPEMBlock)

	if err != nil {
		return err
	}

	mch.wcert[cert.SerialNumber()] = cert

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
