package pkg

import (
	"strings"

	"github.com/thbkrkr/vaultd/pkg/aes"
)

func NewVault(filename string, crypter aes.AESCrypter) Vault {
	// JSON
	if strings.HasSuffix(filename, ".json") {
		return JsonVault{Crypter: crypter}
	}
	// YAML
	if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		return YamlVault{Crypter: crypter}
	}
	// Environment variables (A=B)
	if strings.HasSuffix(filename, ".env") {
		return EnvVault{Crypter: crypter}
	}
	// Else
	return RawVault{Crypter: crypter}
}

type Vault interface {
	Decrypt(data []byte) ([]byte, error)
	Encrypt(data []byte) ([]byte, error)
}
