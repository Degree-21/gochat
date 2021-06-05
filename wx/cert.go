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

	// VerifyWithSha256 签名验证
	VerifyWithSha256(data, signature []byte) error

	// IsValid 证书是否有效
	IsValid() bool

	// ExpiredAt 证书过期时间
	ExpiredAt() time.Time
}

type rsacert struct {
	cert *x509.Certificate
}

func (r *rsacert) SerialNumber() string {
	return fmt.Sprintf("%X", r.cert.SerialNumber)
}

func (r *rsacert) EncryptOEAP(plainText []byte) ([]byte, error) {
	publicKey, ok := r.cert.PublicKey.(*rsa.PublicKey)

	if !ok {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	return rsa.EncryptOAEP(sha1.New(), rand.Reader, publicKey, plainText, nil)
}

func (r *rsacert) VerifyWithSha256(data, signature []byte) error {
	publicKey, ok := r.cert.PublicKey.(*rsa.PublicKey)

	if !ok {
		return errors.New("gochat: invalid rsa public key")
	}

	hashed := sha256.Sum256(data)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
}

func (r *rsacert) IsValid() bool {
	now := time.Now()

	return now.After(r.cert.NotBefore) && now.Before(r.cert.NotAfter)
}

func (r *rsacert) ExpiredAt() time.Time {
	return r.cert.NotAfter
}

// NewRSACert returns a new rsa cert
func NewRSACert(certPEMBlock []byte) (RSACert, error) {
	// x509 cert
	certDERBlock, _ := pem.Decode(certPEMBlock)

	if certDERBlock == nil {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	cert, err := x509.ParseCertificate(certDERBlock.Bytes)

	if err != nil {
		return nil, err
	}

	return &rsacert{cert: cert}, nil
}

// RSAKey RSA密钥
type RSAKey interface {
	// DecryptOEAP OEAP解密
	DecryptOEAP(cipherText []byte) ([]byte, error)

	// SignWithSha256 签名
	SignWithSha256(data []byte) ([]byte, error)
}

type rsakey struct {
	key *rsa.PrivateKey
}

func (r *rsakey) DecryptOEAP(cipherText []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha1.New(), rand.Reader, r.key, cipherText, nil)
}

func (r *rsakey) SignWithSha256(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)

	signature, err := rsa.SignPKCS1v15(rand.Reader, r.key, crypto.SHA256, h.Sum(nil))

	if err != nil {
		return nil, err
	}

	return signature, nil
}

// NewRSAKey returns a new rsa key
func NewRSAKey(keyPEMBlock []byte) (RSAKey, error) {
	keyDERBlock, _ := pem.Decode(keyPEMBlock)

	if keyDERBlock == nil {
		return nil, errors.New("gochat: invalid rsa private key")
	}

	var (
		key interface{}
		err error
	)

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

	return &rsakey{key: privateKey}, nil
}
