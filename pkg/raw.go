package pkg

import (
	"github.com/thbkrkr/vaultd/pkg/aes"
)

type RawVault struct {
	Crypter aes.AESCrypter
}

func (ev RawVault) Decrypt(data []byte) ([]byte, error) {
	str, err := ev.Crypter.Decrypt(string(data))
	return []byte(str), err
}

func (ev RawVault) Encrypt(data []byte) ([]byte, error) {
	str, err := ev.Crypter.Encrypt(string(data))
	return []byte(str), err
}
