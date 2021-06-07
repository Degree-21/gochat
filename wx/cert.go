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
	"os"
	"path/filepath"
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
	sn   string
	cert *x509.Certificate
}

func (r *rsacert) SerialNumber() string {
	return r.sn
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
func NewRSACert(certBlock []byte) (RSACert, error) {
	// x509 cert
	derBlock, _ := pem.Decode(certBlock)

	if derBlock == nil {
		return nil, errors.New("gochat: invalid rsa public key")
	}

	cert, err := x509.ParseCertificate(derBlock.Bytes)

	if err != nil {
		return nil, err
	}

	return &rsacert{
		sn:   fmt.Sprintf("%X", cert.SerialNumber),
		cert: cert,
	}, nil
}

// NewRSACertFromFile returns a new rsa key from the given file
func NewRSACertFromFile(certFile string) (RSACert, error) {
	certPath, err := filepath.Abs(filepath.Clean(certFile))

	if err != nil {
		return nil, err
	}

	certBlock, err := os.ReadFile(certPath)

	if err != nil {
		return nil, err
	}

	return NewRSACert(certBlock)
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
func NewRSAKey(keyBlock []byte) (RSAKey, error) {
	derBlock, _ := pem.Decode(keyBlock)

	if derBlock == nil {
		return nil, errors.New("gochat: invalid rsa private key")
	}

	var (
		key interface{}
		err error
	)

	switch PemBlockType(derBlock.Type) {
	case RSAPKCS1:
		key, err = x509.ParsePKCS1PrivateKey(derBlock.Bytes)
	case RSAPKCS8:
		key, err = x509.ParsePKCS8PrivateKey(derBlock.Bytes)
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

// NewRSAKeyFromFile returns a new rsa key from the given file
func NewRSAKeyFromFile(keyFile string) (RSAKey, error) {
	keyPath, err := filepath.Abs(filepath.Clean(keyFile))

	if err != nil {
		return nil, err
	}

	keyBlock, err := os.ReadFile(keyPath)

	if err != nil {
		return nil, err
	}

	return NewRSAKey(keyBlock)
}
