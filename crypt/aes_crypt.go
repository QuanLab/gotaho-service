package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/QuanLab/gotaho-service/config"
	"io"
	"log"
)

var (
	passphrase string
)

func CreateHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func init() {
	passphrase = config.Get().Server.PrivatePassPhase
}

func Encrypt(data []byte) string {
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func Decrypt(cypherText string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cypherText)
	if err != nil {
		log.Println("Error decode string base64")
	}
	key := []byte(CreateHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext), err
}
