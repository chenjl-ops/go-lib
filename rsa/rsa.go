package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
)

// NewGenerateRSAKey 生成RSA公私钥
func NewGenerateRSAKey() (priKey string, pubKey string, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)

	privateKeyBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privateKeyFile, err := os.Create("/tmp/private.pem")
	if err != nil {
		return "", "", err
	}
	err = pem.Encode(privateKeyFile, &privateKeyBlock)
	if err != nil {
		return "", "", err
	}
	privateKeyFile.Close()

	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", "", err
	}

	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyFile, err := os.Create("/tmp/public.pem")
	if err != nil {
		return "", "", err
	}
	err = pem.Encode(publicKeyFile, &publicKeyBlock)
	if err != nil {
		return "", "", err
	}
	publicKeyFile.Close()
	return base64.StdEncoding.EncodeToString(privateKeyBlock.Bytes), base64.StdEncoding.EncodeToString(publicKeyBlock.Bytes), nil
}

// LoadPrivateKeyFromString 通过字符串获取私钥
func LoadPrivateKeyFromString(privateKey string) (*rsa.PrivateKey, error) {
	if !strings.HasPrefix(privateKey, "-----BEGIN") {
		privateKey = fmt.Sprintf("-----BEGIN RSA PRIVATE KEY-----\n%s", privateKey)
	}
	if !strings.HasSuffix(privateKey, "-----END RSA PRIVATE KEY-----") {
		privateKey = fmt.Sprintf("%s\n-----END RSA PRIVATE KEY-----", privateKey)
	}
	//fmt.Println("privateKeys: ", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadPrivateKey 通过文件获取私钥
func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadPublicKeyFromString 通过字符串获取公钥
func LoadPublicKeyFromString(publicKey string) (*rsa.PublicKey, error) {
	if !strings.HasPrefix(publicKey, "-----BEGIN") {
		publicKey = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s", publicKey)
	}
	if !strings.HasSuffix(publicKey, "-----END PUBLIC KEY-----") {
		publicKey = fmt.Sprintf("%s\n-----END PUBLIC KEY-----", publicKey)
	}
	block, _ := pem.Decode([]byte(publicKey))
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

// LoadPublicKey 通过文件获取公钥
func LoadPublicKey(path string) (*rsa.PublicKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(data)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

// SignRSA 通过私钥签名
func SignRSA(message string, privateKey *rsa.PrivateKey) (string, error) {
	hash := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil

}

// VerifyRSA 通过公钥验证签名
func VerifyRSA(publicKey *rsa.PublicKey, message string, signature string) bool {
	hash := sha256.Sum256([]byte(message))
	signBytes, _ := base64.StdEncoding.DecodeString(signature)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signBytes)
	return err == nil
}
