package main

type RawVault struct {
}

func (ev RawVault) Decrypt(data []byte) ([]byte, error) {
	str, err := AESDecrypt(string(data))
	return []byte(str), err
}

func (ev RawVault) Encrypt(data []byte) ([]byte, error) {
	str, err := AESEncrypt(string(data))
	return []byte(str), err
}
