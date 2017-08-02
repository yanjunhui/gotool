package gotool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

//生成key
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "私钥",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "公钥",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

//Rsa加密
func RsaEncrypt(origData []byte, publicKeyFile string) (string, error) {
	publicKeyByte, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		return "", errors.New("公钥文件打开失败")
	}

	block, _ := pem.Decode(publicKeyByte)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	pubByte, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)

	return base64.StdEncoding.EncodeToString(pubByte), err
}

//Rsa解密
func RsaDecrypt(ciphertext string, privateKey string) ([]byte, error) {
	privateKeyByte, err := ioutil.ReadFile(privateKey)
	if err != nil {
		return nil, errors.New("私钥文件打开失败")
	}

	block, _ := pem.Decode(privateKeyByte)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, decodeBytes)
}
