package main

type EnvVault struct {
}

func (ev EnvVault) Decrypt(data []byte) ([]byte, error) {
	str, err := AESDecrypt(string(data))
	return []byte(str), err
}

func (ev EnvVault) Encrypt(data []byte) ([]byte, error) {
	str, err := AESEncrypt(string(data))
	return []byte(str), err
}
