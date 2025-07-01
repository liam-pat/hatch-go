package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"github.com/forgoer/openssl"
	"log"
)

// (Data Encryption Standard) Can be brute-forced

// key len: 8 byte (56 bit actually)
// if scr len > 64 bit , need to encrypt many times
func main() {

	src := []byte("test the des cbc encryption")
	key := []byte("12345678"[:8])
	iv := []byte("43218765"[:8])

	dst := desCBSEncrypt(src, key, iv)
	decryptedStr, _ := openssl.DesCBCDecrypt(dst, key, iv, openssl.PKCS5_PADDING)

	log.Printf("`%s` encrypted `%v` \n", src, string(dst))
	log.Printf("`%s` encrypted to str `%v` \n", src, hex.EncodeToString(dst))
	log.Printf("`%s` decrypted `%v` \n", string(dst), string(decryptedStr))

	// 3 times DES CBC , key must be 24 byte
	src = []byte("3 DES CBC encryption")
	key = []byte("1234567812345678123456781234567812345678"[:24])

	dst, _ = openssl.Des3ECBEncrypt(src, key, openssl.ZEROS_PADDING)
	decryptedStr, _ = openssl.Des3ECBDecrypt(dst, key, openssl.ZEROS_PADDING)

	log.Printf("`%s` encrypted `%s` \n", src, string(dst))
	log.Printf("`%s` encrypted 2 hex `%v` \n", src, hex.EncodeToString(dst))
	log.Printf("`%s` decrypted 2 str `%v` \n", string(dst), string(decryptedStr))
}

func desCBSEncrypt(src, key, iv []byte) []byte {
	block, _ := des.NewCipher(key)
	blockSize := block.BlockSize()

	need2padLen := blockSize - len(src)%blockSize

	padStr := bytes.Repeat([]byte{byte(need2padLen)}, need2padLen)
	str := append(src, padStr...)

	dst := make([]byte, len(str))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(dst, str)
	return dst
}
