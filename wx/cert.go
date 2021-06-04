package wx

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

// RSACert RSA证书
type RSACert interface {
	// SerialNumber 证书序列号
	SerialNumber() string

	// EncryptOEAP OEAP加密
	EncryptOEAP(plainText []byte) ([]byte, error)

	// DecryptOEAP OEAP解密
	DecryptOEAP(cipherText []byte) ([]byte, error)

	// SignWithSha256 签名
	SignWithSha256(data []byte) ([]byte, error)

	// VerifyWithSha256 签名验证
	VerifyWithSha256(data, signature []byte) error

	// IsInvalid 证书是否有效
	IsInvalid() bool

	// IsExpired 证书是否过期
	IsExpired() bool
}

type rsacert struct {
	serialNumber string
	publicKey    *rsa.PublicKey
	privateKey   *rsa.PrivateKey
	notBefore    time.Time
	notAfter     time.Time
}

func (cert *rsacert) SerialNumber() string {
	return cert.serialNumber
}

func (cert *rsacert) EncryptOEAP(plainText []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha1.New(), rand.Reader, cert.publicKey, plainText, nil)
}

func (cert *rsacert) DecryptOEAP(cipherText []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, cert.privateKey, cipherText, nil)
}

func (cert *rsacert) SignWithSha256(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)

	signature, err := rsa.SignPKCS1v15(rand.Reader, cert.privateKey, crypto.SHA256, h.Sum(nil))

	if err != nil {
		return nil, err
	}

	return signature, nil
}

func (cert *rsacert) VerifyWithSha256(data, signature []byte) error {
	hashed := sha256.Sum256(data)

	return rsa.VerifyPKCS1v15(cert.publicKey, crypto.SHA256, hashed[:], signature)
}

func (cert *rsacert) IsInvalid() bool {
	now := time.Now()

	return now.After(cert.notBefore) && now.Before(cert.notAfter)
}

func (cert *rsacert) IsExpired() bool {
	return time.Now().After(cert.notAfter)
}

func NewRSACert(certPEMBlock, keyPEMBlock []byte) (RSACert, error) {
	// x509 cert
	certDERBlock, _ := pem.Decode(certPEMBlock)

	if certDERBlock == nil {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)

	if err != nil {
		return nil, err
	}

	publicKey, ok := x509Cert.PublicKey.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	keyDERBlock, _ := pem.Decode(keyPEMBlock)

	if certDERBlock == nil {
		return nil, errors.New("gochat: invalid rsa private key")
	}

	// private key
	var key interface{}

	switch PemBlockType(keyDERBlock.Type) {
	case RSAPKCS1:
		key, err = x509.ParsePKCS1PrivateKey(keyDERBlock.Bytes)
	case RSAPKCS8:
		key, err = x509.ParsePKCS8PrivateKey(keyDERBlock.Bytes)
	}

	if err != nil {
		return nil, err
	}

	privateKey, ok := key.(*rsa.PrivateKey)

	if !ok {
		return nil, errors.New("gochat: invalid rsa private key")
	}

	return &rsacert{
		serialNumber: fmt.Sprintf("%X", x509Cert.SerialNumber),
		publicKey:    publicKey,
		privateKey:   privateKey,
		notBefore:    x509Cert.NotBefore,
		notAfter:     x509Cert.NotAfter,
	}, nil
}
