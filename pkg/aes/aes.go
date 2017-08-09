package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

type AESCrypter struct {
	cipherKey []byte
	nonce     []byte
}

func NewAESCrypter(key string, nonce string) AESCrypter {
	bnonce := hex.EncodeToString([]byte(nonce))
	snonce, _ := hex.DecodeString(bnonce)
	return AESCrypter{
		cipherKey: []byte(key),
		nonce:     snonce,
	}
}

func (c *AESCrypter) Encrypt(plaintext string) (decodedmess string, err error) {
	block, err := aes.NewCipher(c.cipherKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	cipherText := aesgcm.Seal(nil, c.nonce, []byte(plaintext), nil)
	return fmt.Sprintf("%x", cipherText), nil
}

func (c *AESCrypter) Decrypt(message string) (string, error) {
	cipherText, err := hex.DecodeString(message)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(c.cipherKey)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, c.nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", string(plaintext)), nil
}
