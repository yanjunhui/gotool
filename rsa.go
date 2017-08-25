package gotool

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
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
func RsaEncrypt(origData []byte, publicKey []byte) (encryptStr string, err error) {
	var maxSize = 117
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	ecrypt := func(data []byte) (string, error) {
		eData, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
		if err != nil {
			return "", err
		}
		return base64.StdEncoding.EncodeToString(eData), err
	}

	for len(origData) > 0 {
		if len(origData) > maxSize {
			tempData, err := ecrypt(origData[:maxSize])
			if err != nil {
				return "", err
			}
			encryptStr += tempData
			origData = origData[maxSize:]
		} else {
			tempData, err := ecrypt(origData)
			if err != nil {
				return "", err
			}
			encryptStr += tempData
			origData = nil
		}
	}

	return encryptStr, nil
}

//Rsa解密
func RsaDecrypt(ciphertext string, privateKey []byte) (decruptData []byte, err error) {
	var maxSize = 172
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	decode := func(str string) ([]byte, error) {
		decodeBytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			return nil, err
		}
		return rsa.DecryptPKCS1v15(rand.Reader, priv, decodeBytes)
	}

	for n := len(ciphertext) / maxSize; n > 0; n-- {
		dData, err := decode(ciphertext[:maxSize])
		if err != nil {
			return nil, err
		}

		decruptData = bytes.Join([][]byte{decruptData, dData}, []byte(""))
		ciphertext = ciphertext[maxSize:]
	}

	return decruptData, nil

}
