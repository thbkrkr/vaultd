package aes

import (
	"fmt"
	"log"
	"testing"
)

var c AESCrypter

func init() {
	c = AESCrypter{
		CipherKey: []byte("wpunIGR08kMX6pI8gWPBrFQApwcFXbpR"),
		Nonce:     "afb8a7579bf971db9f8ceeed",
		//Sault:     "7XDW52iEB580QAvp",
	}
}

func TestAES(t *testing.T) {
	originalMsg := "42"

	encryptedMsg, err := c.AESEncrypt(originalMsg)
	if err != nil {
		log.Fatal(err)
	}

	decryptedMsg, err := c.AESDecrypt(encryptedMsg)
	if err != nil {
		log.Fatal(err)
	}

	if decryptedMsg != originalMsg {
		log.Fatal(fmt.Errorf(
			"message expected: '%s', equals: '%s'", originalMsg, decryptedMsg))
	}
}
