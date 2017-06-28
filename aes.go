package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/Sirupsen/logrus"
)

var CipherKey = []byte("wpunIGR08kMX6pI8gWPBrFQApwcFXbpR")
var Nonce = "afb8a7579bf971db9f8ceeed"
var Sault = "7XDW52iEB580QAvp"

func AESDecrypt(text string) (string, error) {
	block, err := aes.NewCipher(CipherKey)
	if err != nil {
		logrus.Error("Error initializing cypher")
		return "", err
	}

	cipherText, err := hex.DecodeString(text)
	if err != nil {
		return "", err
	}

	/*
		cipherText, err := base64.StdEncoding.DecodeString(text)
		if err != nil {
			logrus.Error("Error decoding base64 encrypted data")
			return "", err
		}
	*/

	if len(cipherText) < aes.BlockSize {
		logrus.Errorf("ciphertext too short: %d", len(cipherText))
		return "", errors.New(fmt.Sprint())
	}
	unb64ed := []byte(cipherText)
	iv := unb64ed[:aes.BlockSize]
	unb64ed = unb64ed[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(unb64ed, unb64ed)
	return string(unb64ed), nil
}

func AESdecrypt2(message string) (string, error) {
	cipherText, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}

	nonce, _ := hex.DecodeString(Nonce)

	block, err := aes.NewCipher(CipherKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", string(plaintext)), nil
	//return base64.URLEncoding.EncodeToString(plaintext), nil
}

func AESEncrypt(text string) (string, error) {
	Salt := fmt.Sprintf("%d", time.Now().Nanosecond())
	block, err := aes.NewCipher(CipherKey)
	if err != nil {
		logrus.Error("Error initializing cypher")
		return "", err
	}
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	si := 0
	for i, _ := range iv {
		iv[i] = Salt[si]
		si = (si + 1) % len(Salt)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(text))

	return fmt.Sprintf("%x", cipherText), nil
	//return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AESencrypt2(plaintext string) (decodedmess string, err error) {
	/*cipherText, err := base64.URLEncoding.DecodeString(securemess)
		if err != nil {
	    return "", err
	  }*/

	block, err := aes.NewCipher(CipherKey)
	if err != nil {
		return "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	cipherText := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return fmt.Sprintf("%x", cipherText), nil
}
