package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/thbkrkr/go-utilz/http"
)

var (
	name      = "vaultd"
	buildDate = "dev"
	gitCommit = "dev"
	port      = 4242
)

func main() {
	http.API(name, buildDate, gitCommit, port, func(r *gin.Engine) {

		r.GET("/s/*name", func(c *gin.Context) {
			name := c.Param("name")
			mode := c.Query("mode")

			secret, err := lookupSecretFile(mode, name)
			if err != nil {
				log.WithError(err).WithField("name", name).Error("Fail to lookup secret")
				c.JSON(400, "Secret lookup failed")
				return
			}

			if strings.HasSuffix(name, ".json") {
				c.Data(200, "application/json; charset=utf-8", secret)
			} else if strings.HasSuffix(name, ".yaml") {
				c.Data(200, "application/x-yaml; charset=utf-8", secret)
			} else {
				c.Data(200, "text/plain; charset=utf-8", secret)
			}
		})

		r.GET("/list/*path", func(c *gin.Context) {
			rootPath := c.Param("path")
			secrets := []string{}

			err := filepath.Walk("."+rootPath, func(secretPath string, f os.FileInfo, err error) error {
				if "."+rootPath == secretPath {
					return nil
				}
				secrets = append(secrets, secretPath)
				return nil
			})

			if err != nil {
				log.WithError(err).WithField("path", rootPath).Error("Fail to list secrets")
				c.JSON(400, "Fail to list secrets")
				return
			}

			c.JSON(200, secrets)
		})
	})
}

func getVault(filename string) Vault {
	if strings.HasSuffix(filename, ".json") {
		return JsonVault{}
	} else if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		return YamlVault{}
	} else if strings.HasSuffix(filename, ".env") {
		return EnvVault{}
	} else {
		return RawVault{}
	}
}

func lookupSecretFile(mode string, name string) ([]byte, error) {
	suffix := ".encrypt"
	if mode == "encrypt" {
		suffix = ""
	}

	fmt.Println("." + name + suffix)

	rawData, err := ioutil.ReadFile("." + name + suffix)
	if err != nil {
		return nil, err
	}

	vault := getVault(name)

	var data []byte
	if mode == "encrypt" {
		data, err = vault.Encrypt(rawData)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = vault.Decrypt(data)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

type Vault interface {
	Decrypt(data []byte) ([]byte, error)
	Encrypt(data []byte) ([]byte, error)
}

func decryptVal(data string) (string, error) {
	return AESDecrypt(data)
}

func encryptVal(data string) (string, error) {
	return AESEncrypt(data)
}
