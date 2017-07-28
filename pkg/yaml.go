package pkg

import (
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/thbkrkr/vaultd/pkg/aes"
	"gopkg.in/yaml.v2"
)

type YamlVault struct {
	Crypter aes.AESCrypter
}

func (yv YamlVault) Decrypt(data []byte) ([]byte, error) {
	yml := make(map[string]interface{})
	yml2 := make(map[string]interface{})
	err := yaml.Unmarshal(data, &yml)
	if err != nil {
		return nil, err
	}

	for k, v := range yml {
		yml2[k] = yv.decryptVal(v)
	}

	bytes, err := yaml.Marshal(yml2)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (yv YamlVault) decryptVal(v interface{}) interface{} {
	var v2 interface{}
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice {
		v2 = yv.decryptSlice(v.([]interface{}))
	} else if typ == reflect.Map {
		v2 = yv.decryptMap(v.(map[interface{}]interface{}))
	} else {
		v2, _ = yv.Crypter.AESDecrypt(fmt.Sprintf("%v", v))
	}
	return v2
}

func (yv YamlVault) decryptMap(m map[interface{}]interface{}) map[interface{}]interface{} {
	m2 := make(map[interface{}]interface{}, len(m))
	for k, v := range m {
		m2[k] = yv.decryptVal(v)
	}
	return m2
}

func (yv YamlVault) decryptSlice(slc []interface{}) []interface{} {
	slc2 := make([]interface{}, len(slc))
	for i, v := range slc {
		slc2[i] = yv.decryptVal(v)
	}
	return slc2
}

func (yv YamlVault) Encrypt(data []byte) ([]byte, error) {
	yml := make(map[string]interface{})
	yml2 := make(map[string]interface{})
	err := yaml.Unmarshal(data, &yml)
	if err != nil {
		return nil, err
	}

	for k, v := range yml {
		yml2[k] = yv.encryptVal(v)
	}

	bytes, err := yaml.Marshal(yml2)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (yv YamlVault) encryptVal(v interface{}) interface{} {
	var v2 interface{}
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice {
		v2 = yv.encryptSlice(v.([]interface{}))
	} else if typ == reflect.Map {
		v2 = yv.encryptMap(v.(map[interface{}]interface{}))
	} else {
		var err error
		v2, err = yv.Crypter.AESEncrypt(fmt.Sprintf("%v", v))
		if err != nil {
			logrus.WithError(err).WithField("val", v).Error("Fail to encrypt val")
		}
	}
	return v2
}

func (yv YamlVault) encryptMap(m map[interface{}]interface{}) map[interface{}]interface{} {
	m2 := make(map[interface{}]interface{}, len(m))
	for k, v := range m {
		m2[k] = yv.encryptVal(v)
	}
	return m2
}

func (yv YamlVault) encryptSlice(slc []interface{}) []interface{} {
	slc2 := make([]interface{}, len(slc))
	for i, v := range slc {
		slc2[i] = yv.encryptVal(v)
	}
	return slc2
}
