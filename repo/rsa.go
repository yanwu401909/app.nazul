package repo

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	. "app.nazul/errors"
	"app.nazul/models"
)

func GenerateRsaKey(bits int) ApiError {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return NewErrorWithMessage(SERVER_ERROR, err.Error())
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	private := base64.StdEncoding.EncodeToString(block.Bytes)

	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return NewErrorWithMessage(SERVER_ERROR, err.Error())
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	public := base64.StdEncoding.EncodeToString(block.Bytes)
	if err := CONN.Create(&models.RsaPair{
		PublicKey:  public,
		PrivateKey: private,
	}).Error; err != nil {
		return NewErrorWithMessage(DB_ERROR, err.Error())
	}
	return ApiError{}
}

func RandomRsaPair() (data models.RsaPair, errs ApiError) {
	if result := CONN.Where("id = ?", 1).First(&data); result.Error != nil {
		errs = NewErrorWithMessage(DB_ERROR, result.Error.Error())
	}
	return
}
