package pkg

import (
	"github.com/thbkrkr/vaultd/pkg/aes"
)

type RawVault struct {
	Crypter aes.AESCrypter
}

func (ev RawVault) Decrypt(data []byte) ([]byte, error) {
	str, err := ev.Crypter.AESDecrypt(string(data))
	return []byte(str), err
}

func (ev RawVault) Encrypt(data []byte) ([]byte, error) {
	str, err := ev.Crypter.AESEncrypt(string(data))
	return []byte(str), err
}
