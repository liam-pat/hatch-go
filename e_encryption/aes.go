package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
	"strings"
)

// Advanced Encryption Standard
// @todo, study 3 method 's different
func main() {

	str := "hello world!!! I am the AI"
	// the key length must be 16,24,32
	key := "123456781234567812345678"

	encryptedStr, _ := aesCBCEncrypt([]byte(str), []byte(key))
	decryptedStr, _ := aesCBCDecrypt(encryptedStr, []byte(key))
	log.Printf("uses CBCEncrypter encrypt `%s` to `%s`\n", str, string(encryptedStr))
	log.Printf("uses CBCDecrypter decrypt `%s` to `%s`\n", str, decryptedStr)
	log.Println(strings.Repeat("#", 20))

	encryptedStr = aesECBEncrypt([]byte(str), []byte(key))
	decryptedStr = aesECBDecrypt(encryptedStr, []byte(key))
	log.Printf("uses ECBEncrypter encrypt `%s` to `%s`\n", str, string(encryptedStr))
	log.Printf("uses ECBDecrypter decrypt `%s` to `%s`\n", str, decryptedStr)
	log.Println(strings.Repeat("#", 20))

	cfbEncryptedStr := aesCFBEncrypt([]byte(str), []byte(key))
	cfbEncryptedHexStr := hex.EncodeToString(cfbEncryptedStr)
	cfbDecryptedStr := aesCFBDecrypt(cfbEncryptedStr, []byte(key))
	log.Printf("uses CFBEncrypter encrypt `%s` to `%s`\n", str, string(cfbEncryptedStr))
	log.Printf("uses CFBEncrypter encrypt `%s` to hex string `%s`\n", str, cfbEncryptedHexStr)
	log.Printf("uses CFBDecrypter decrypt `%s` to `%s`\n", str, cfbDecryptedStr)
	log.Println(strings.Repeat("#", 20))
}

func aesCBCEncrypt(src, key []byte) ([]byte, error) {
	keyBlock, _ := aes.NewCipher(key)
	keyBlockLen := keyBlock.BlockSize()

	padLen := keyBlockLen - len(src)%keyBlockLen
	padStr := bytes.Repeat([]byte{byte(padLen)}, padLen)
	originalStrPadded := append(src, padStr...)

	blockMode := cipher.NewCBCEncrypter(keyBlock, key[:keyBlockLen])
	dst := make([]byte, len(originalStrPadded))
	blockMode.CryptBlocks(dst, originalStrPadded)

	return dst, nil
}

func aesCBCDecrypt(dst, key []byte) ([]byte, error) {
	keyBlock, _ := aes.NewCipher(key)
	keyBlockLen := keyBlock.BlockSize()
	blockMode := cipher.NewCBCDecrypter(keyBlock, key[:keyBlockLen])

	scr := make([]byte, len(dst))
	blockMode.CryptBlocks(scr, dst)

	scrLength := len(scr)
	paddingNum := int(scr[scrLength-1])

	return scr[:(scrLength - paddingNum)], nil
}

// @todo need to simplify
func aesECBEncrypt(scr, key []byte) []byte {
	newKey := make([]byte, 16)
	copy(newKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			newKey[j] ^= key[i]
		}
	}

	keyBlock, _ := aes.NewCipher(newKey)
	length := (len(scr) + aes.BlockSize) / aes.BlockSize

	plain := make([]byte, length*aes.BlockSize)
	copy(plain, scr)
	pad := byte(len(plain) - len(scr))
	for i := len(scr); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, keyBlock.BlockSize(); bs <= len(scr); bs, be = bs+keyBlock.BlockSize(), be+keyBlock.BlockSize() {
		keyBlock.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// @todo need to simplify
func aesECBDecrypt(dst, key []byte) (scr []byte) {
	newKey := make([]byte, 16)
	copy(newKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			newKey[j] ^= key[i]
		}
	}

	keyBlock, _ := aes.NewCipher(newKey)
	decryptedFullStr := make([]byte, len(dst))
	for bs, be := 0, keyBlock.BlockSize(); bs < len(dst); bs, be = bs+keyBlock.BlockSize(), be+keyBlock.BlockSize() {
		keyBlock.Decrypt(decryptedFullStr[bs:be], dst[bs:be])
	}

	trim := 0
	if len(decryptedFullStr) > 0 {
		trim = len(decryptedFullStr) - int(decryptedFullStr[len(decryptedFullStr)-1])
	}

	return decryptedFullStr[:trim]
}

func aesCFBEncrypt(scr []byte, key []byte) (encrypted []byte) {
	block, _ := aes.NewCipher(key)

	encrypted = make([]byte, aes.BlockSize+len(scr))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], scr)
	return encrypted
}

func aesCFBDecrypt(dst []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)

	iv := dst[:aes.BlockSize]
	dst = dst[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(dst, dst)
	return dst
}
