package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func main() {

	b := []byte("world")
	str, err := RsaWithSHA256Base64("hello", b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

func RsaWithSHA256Base64(origData string, block []byte) (sign string, err error) {
	//blocks, _ := pem.Decode(block)
	privateKey, _ := x509.ParsePKCS8PrivateKey(block)
	h := sha256.New()
	h.Write([]byte(origData))
	digest := h.Sum(nil)
	s, _ := rsa.SignPKCS1v15(nil, privateKey.(*rsa.PrivateKey), crypto.SHA256, digest)
	sign = base64.StdEncoding.EncodeToString(s)
	return
}
