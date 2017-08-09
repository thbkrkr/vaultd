package aes

import (
	"fmt"
	"log"
	"testing"
)

var c AESCrypter

func init() {
	c = NewAESCrypter("AES256Key-32Characters1234567890", "nonce-12Char")
}

func TestAES(t *testing.T) {
	originalMsg := "42"

	encryptedMsg, err := c.Encrypt(originalMsg)
	if err != nil {
		log.Fatal("Fail to encrypt: ", err.Error())
	}

	decryptedMsg, err := c.Decrypt(encryptedMsg)
	if err != nil {
		log.Fatal("Fail to decrypt: ", err.Error())
	}

	if decryptedMsg != originalMsg {
		log.Fatal(fmt.Errorf(
			"message expected: '%s', equals: '%s'", originalMsg, decryptedMsg))
	}
}
