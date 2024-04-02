package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func main() {
	origText := "hello world"

	key := "123456781234567812345678"
	fmt.Println("original text：", origText)

	encryptCode := aesEncrypt(origText, key)
	fmt.Println("encrypted text：", encryptCode)

	decryptCode := aesDecrypt(encryptCode, key)
	fmt.Println("decrypted text：", decryptCode)
}

func aesEncrypt(origText string, key string) string {
	byteData := []byte(origText)
	byteKey := []byte(key)

	block, _ := aes.NewCipher(byteKey)
	blockSize := block.BlockSize()

	// padding , data length = key block len
	paddingLen := blockSize - len(byteData)%blockSize
	padText := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	byteData = append(byteData, padText...)
	fmt.Printf("key length: %d \t block size %d \t data length: %d\n", len(byteKey), blockSize, len(byteData))

	blockMode := cipher.NewCBCEncrypter(block, byteKey[:blockSize])
	emptyByte := make([]byte, len(byteData))
	blockMode.CryptBlocks(emptyByte, byteData)

	return base64.StdEncoding.EncodeToString(emptyByte)
}

func aesDecrypt(str string, key string) string {
	encryptedByte, _ := base64.StdEncoding.DecodeString(str)
	k := []byte(key)

	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	emptyByte := make([]byte, len(encryptedByte))
	blockMode.CryptBlocks(emptyByte, encryptedByte)

	unPaddingLen := int(emptyByte[len(emptyByte)-1])
	originalByte := emptyByte[:(len(emptyByte) - unPaddingLen)]

	return string(originalByte)
}
