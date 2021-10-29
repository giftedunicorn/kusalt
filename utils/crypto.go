package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/mergermarket/go-pkcs7"
	log "github.com/sirupsen/logrus"
	"io"
	"strings"
)

const CIPHER_KEY = "w!z%C*F)J@NcRfUjXn2r5u8x/A?D(G+K"

// https://gist.github.com/brettscott/2ac58ab7cb1c66e2b4a32d6c1c3908a7
// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string) (string, error) {
	log.Info("Decrypt invoked")
	key := []byte(CIPHER_KEY)

	ivStr := strings.Split(encrypted, ":")[0]
	cipherTextStr := strings.Split(encrypted, ":")[1]
	iv, _ := hex.DecodeString(ivStr)
	cipherText, _ := hex.DecodeString(cipherTextStr)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	return fmt.Sprintf("%x", cipherText), nil
}

// Encrypt encrypts plain text string into cipher text string
func Encrypt(unencrypted string) (string, error) {
	log.Info("Decrypt invoked")
	key := []byte(CIPHER_KEY)
	plainText := []byte(unencrypted)

	plainText, err := pkcs7.Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}
