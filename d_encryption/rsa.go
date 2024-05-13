package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"log"
	"os"
	"strings"
)

func main() {

	str, _ := os.Getwd()

	genRSAKey(1024, str+"/tmp")
	privateKey, publicKey := getRSA(str + "/tmp")

	privateBlock, _ := pem.Decode(privateKey)
	publicBlock, _ := pem.Decode(publicKey)

	anyPrivateKey, _ := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
	anyPublicKey, _ := x509.ParsePKIXPublicKey(publicBlock.Bytes)

	json := "for testing the RSA encryption"
	dst := rsaEncrypt([]byte(json), anyPublicKey)
	decryptedStr := rsaDecrypt(dst, anyPrivateKey)

	log.Printf(" `%s` encrypted 2 hex string `%s`\n", json, hex.EncodeToString(dst))
	log.Printf(strings.Repeat("##", 50))
	log.Printf(" `%s` decrypted 2 `%s`\n", hex.EncodeToString(dst), string(decryptedStr))

}

func genRSAKey(bits int, dir string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	publicKey := privateKey.PublicKey

	// most popular ways : PKCS1 and PKCS8
	X509PrivateKey, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	X509PublicKey, _ := x509.MarshalPKIXPublicKey(&publicKey)

	privateKeyFile, _ := os.Create(dir + "/private.pem")
	publicKeyFile, _ := os.Create(dir + "/public.pem")

	defer func(privateKeyFile *os.File) { privateKeyFile.Close() }(privateKeyFile)
	defer func(publicKeyFile *os.File) { publicKeyFile.Close() }(publicKeyFile)

	privateBlock := pem.Block{Type: "PRIVATE KEY", Bytes: X509PrivateKey}
	publicBlock := pem.Block{Type: "Public Key", Bytes: X509PublicKey}

	_ = pem.Encode(privateKeyFile, &privateBlock)
	_ = pem.Encode(publicKeyFile, &publicBlock)
}

func getRSA(dir string) (privateKeyStr, publicKeyStr []byte) {
	privateKeyFile, _ := os.Open(dir + "/private.pem")
	publicKeyFile, _ := os.Open(dir + "/public.pem")

	defer func(privateKeyFile *os.File) { privateKeyFile.Close() }(privateKeyFile)
	defer func(publicKeyFile *os.File) { publicKeyFile.Close() }(publicKeyFile)

	privateInfo, _ := privateKeyFile.Stat()
	publicInfo, _ := publicKeyFile.Stat()

	privateBuf := make([]byte, privateInfo.Size())
	publicBuf := make([]byte, publicInfo.Size())

	_, _ = privateKeyFile.Read(privateBuf)
	_, _ = publicKeyFile.Read(publicBuf)

	return privateBuf, publicBuf
}

func rsaEncrypt(src []byte, anyPublicKey any) (dst []byte) {
	dst, _ = rsa.EncryptPKCS1v15(rand.Reader, anyPublicKey.(*rsa.PublicKey), src)

	return
}

func rsaDecrypt(dst []byte, anyPrivateKey any) (src []byte) {
	src, _ = rsa.DecryptPKCS1v15(rand.Reader, anyPrivateKey.(*rsa.PrivateKey), dst)
	return
}
