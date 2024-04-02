package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
)

func main() {
	key := []byte("2fa6c1e9")
	str := "I love this beautiful world!"

	strEncrypted, _ := encrypt(str, key)
	fmt.Println("Encrypted:", strEncrypted)

	strDecrypted, _ := decrypt(strEncrypted, key)
	fmt.Println("Decrypted:", strDecrypted)
}

func decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func encrypt(str string, key []byte) (string, error) {
	byteStr := []byte(str)

	block, _ := des.NewCipher(key)
	blockSize := block.BlockSize()

	padding := blockSize - len(byteStr)%blockSize
	padBytes := bytes.Repeat([]byte{0}, padding)
	byteStr = append(byteStr, padBytes...)

	emptyStr := make([]byte, len(byteStr))
	dst := emptyStr

	for len(byteStr) > 0 {
		block.Encrypt(dst, byteStr[:blockSize])
		byteStr = byteStr[blockSize:]
		dst = dst[blockSize:]
	}
	return hex.EncodeToString(emptyStr), nil
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}
