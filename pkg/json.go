package pkg

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/thbkrkr/vaultd/pkg/aes"
)

type JsonVault struct {
	Crypter aes.AESCrypter
}

func (jv JsonVault) Decrypt(data []byte) ([]byte, error) {
	encryptedJson := make(map[string]interface{})
	decryptedJson := make(map[string]interface{})

	err := json.Unmarshal(data, &encryptedJson)
	if err != nil {
		return nil, err
	}

	for k, v := range encryptedJson {
		decryptedJson[k] = jv.decryptVal(v)
	}

	bytes, err := json.Marshal(decryptedJson)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (jv JsonVault) decryptVal(v interface{}) interface{} {
	var v2 interface{}
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice {
		v2 = jv.decryptSlice(v.([]interface{}))
	} else if typ == reflect.Map {
		v2 = jv.decryptMap(v.(map[string]interface{}))
	} else {
		var err error
		v2, err = jv.Crypter.Decrypt(fmt.Sprintf("%v", v))
		if err != nil {
			logrus.WithError(err).WithField("val", v).Error("Fail to decrypt val")
		}
	}
	return v2
}

func (jv JsonVault) decryptMap(m map[string]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{}, len(m))
	for k, v := range m {
		m2[k] = jv.decryptVal(v)
	}
	return m2
}

func (jv JsonVault) decryptSlice(slc []interface{}) []interface{} {
	slc2 := make([]interface{}, len(slc))
	for i, v := range slc {
		slc2[i] = jv.decryptVal(v)
	}
	return slc2
}

func (jv JsonVault) Encrypt(data []byte) ([]byte, error) {
	jso := make(map[string]interface{})
	jso2 := make(map[string]interface{})
	err := json.Unmarshal(data, &jso)
	if err != nil {
		return nil, err
	}

	for k, v := range jso {
		jso2[k] = jv.encryptVal(v)
	}

	bytes, err := json.Marshal(jso2)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (jv JsonVault) encryptVal(v interface{}) interface{} {
	var v2 interface{}
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Slice {
		v2 = jv.encryptSlice(v.([]interface{}))
	} else if typ == reflect.Map {
		v2 = jv.encryptMap(v.(map[string]interface{}))
	} else {
		var err error
		v2, err = jv.Crypter.Encrypt(fmt.Sprintf("%v", v))
		if err != nil {
			logrus.WithError(err).WithField("val", v).Error("Fail to encrypt val")
		}
	}
	return v2
}

func (jv JsonVault) encryptMap(m map[string]interface{}) map[string]interface{} {
	m2 := make(map[string]interface{}, len(m))
	for k, v := range m {
		m2[k] = jv.encryptVal(v)
	}
	return m2
}

func (jv JsonVault) encryptSlice(slc []interface{}) []interface{} {
	slc2 := make([]interface{}, len(slc))
	for i, v := range slc {
		slc2[i] = jv.encryptVal(v)
	}
	return slc2
}
