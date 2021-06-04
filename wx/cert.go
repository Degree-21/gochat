package wx

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type RSACert interface {
	SerialNumber() string
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
	SignWithSha256(data []byte) ([]byte, error)
	VerifyWithSha256(data, signature []byte) error
}

type rsacert struct {
	serialNumber string
	publicKey    *rsa.PublicKey
	privateKey   *rsa.PrivateKey
}

func (cert *rsacert) SerialNumber() string {
	return cert.serialNumber
}

func (cert *rsacert) Encrypt(plainText []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, cert.publicKey, plainText)
}

func (cert *rsacert) Decrypt(cipherText []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, cert.privateKey, cipherText)
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

func NewRSACertFromPemFile(certFile, keyFile string) (RSACert, error) {

	return new(rsacert), nil
}

func NewRSACertFromPemBlock(certPEMBlock, keyPEMBlock []byte) (RSACert, error) {
	// fmt.Sprintf("%X", cert.Leaf.SerialNumber)
	return new(rsacert), nil
}
