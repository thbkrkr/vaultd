package main

import (
	"fmt"
	"log"
	"testing"
)

func TestAES(t *testing.T) {
	message := "42"
	m, err := AESEncrypt(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("m:", m)

	m2, err := AESDecrypt(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("m2:", m)

	if m2 != message {
		fmt.Errorf("message equals: "+m2, "expected: "+message)
	}
}
