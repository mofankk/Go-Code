package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func padding(src []byte, blocksize int) []byte {
	padnum := blocksize - len(src)%blocksize
	pad := bytes.Repeat([]byte{byte(padnum)}, padnum)
	return append(src, pad...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	unpadnum := int(src[n-1])
	return src[:n-unpadnum]
}

func encryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockmode := cipher.NewCBCEncrypter(block, key)
	blockmode.CryptBlocks(src, src)
	return src
}

func decryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockmode := cipher.NewCBCDecrypter(block, key)
	blockmode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}

func main() {
	x := []byte("hello,world")
	key := []byte("hgfedcba87654321")
	x1 := encryptAES(x, key)
	fmt.Println(string(x1))
	x2 := decryptAES(x1, key)
	fmt.Print(string(x2))

	fmt.Println(string(decryptAES([]byte("MTYzMDY2MjczMnxJUXdBSG1OMWFYbHZkWEJsYm1kQU1UUXpNemN5T1RVek16Z3hPRFk1T1RjM05nPT18kr1tOOEsxOKoqUWdjt5BOkQ7Dqfv6kmFPYIc6r5lZuM%3D"), key)))
}
