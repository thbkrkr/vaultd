package pkg

import (
	"bytes"

	"github.com/thbkrkr/vaultd/pkg/aes"
)

type EnvVault struct {
	Crypter aes.AESCrypter
}

func (ev EnvVault) Decrypt(data []byte) ([]byte, error) {
	decryptedEnv := [][]byte{}

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		kv := bytes.Split(line, []byte{'='})
		if len(kv) == 2 {
			v, err := ev.Crypter.AESDecrypt(string(kv[1]))
			if err != nil {
				return nil, err
			}
			line = append(append(kv[0], '='), []byte(v)...)
		}
		decryptedEnv = append(decryptedEnv, line)
	}

	return bytes.Join(decryptedEnv, []byte{'\n'}), nil
}

func (ev EnvVault) Encrypt(data []byte) ([]byte, error) {
	encryptedEnv := []byte{}

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		kv := bytes.Split(line, []byte{'='})
		if len(kv) >= 2 {
			v, err := ev.Crypter.AESEncrypt(string(kv[1]))
			if err != nil {
				return nil, err
			}
			// FIXME: what's the best way to concat []byte and string?
			//line = append(append(kv[0], '='), v...)
			line = []byte(string(append(kv[0], '=')) + v)
		}

		encryptedEnv = append(encryptedEnv, line...)
		encryptedEnv = append(encryptedEnv, '\n')
	}

	return encryptedEnv, nil
}
