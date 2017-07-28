package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/thbkrkr/go-utilz/http"
	"github.com/thbkrkr/vaultd/pkg"
	"github.com/thbkrkr/vaultd/pkg/aes"
)

var (
	name      = "vaultd"
	buildDate = "dev"
	gitCommit = "dev"
	port      = 4242

	dataDir string

	crypter aes.AESCrypter
)

func init() {
	flag.StringVar(&dataDir, "data-dir", ".", "Secrets storage directory")
	flag.Parse()

	dataDir = strings.TrimPrefix(dataDir, "./")
	dataDir = strings.TrimSuffix(dataDir, "/")
}

func main() {
	// TODO: env vars
	crypter = aes.AESCrypter{
		CipherKey: []byte("wpunIGR08kMX6pI8gWPBrFQApwcFXbpR"),
		Nonce:     "afb8a7579bf971db9f8ceeed",
	}

	http.API(name, buildDate, gitCommit, port, func(r *gin.Engine) {
		r.GET("/help", func(c *gin.Context) {
			c.JSON(200, []string{
				"/list/*path      List secrets",
				"/get/*name?mode  Get secrets (default mode: decrypt)",
			})
		})

		r.GET("/ls", func(c *gin.Context) {
			lsFiles(c, "")
		})

		r.GET("/ls/*path", func(c *gin.Context) {
			path := c.Param("path")
			lsFiles(c, path)
		})

		r.GET("/get/*name", func(c *gin.Context) {
			name := c.Param("name")
			mode := c.Query("mode")

			secret, err := GetFile(mode, name)
			if err != nil {
				if strings.Contains(err.Error(), "no such file or directory") {
					c.JSON(400, "Secret `"+name+"` not found")
					return
				}
				log.WithError(err).WithField("name", name).Error("Fail to lookup secret")
				c.JSON(400, "Secret lookup failed")
				return
			}

			format := "text/plain; charset=utf-8"
			if strings.HasSuffix(name, ".json") {
				format = "application/json; charset=utf-8"
			} else if strings.HasSuffix(name, ".yaml") {
				format = "application/x-yaml; charset=utf-8"
			}

			c.Data(200, format, secret)
		})

	})
}

func lsFiles(c *gin.Context, path string) {
	suffix := ".encrypt"
	secrets := []string{}
	dirPath := dataDir + path

	err := filepath.Walk(dirPath, func(secretPath string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(secretPath, suffix) {
			return nil
		}
		if dirPath == secretPath {
			return nil
		}
		secretPath = strings.TrimPrefix(secretPath, dataDir+"/")
		secretPath = strings.TrimSuffix(secretPath, suffix)
		secrets = append(secrets, secretPath)
		return nil
	})

	if err != nil {
		log.WithError(err).WithField("path", path).Error("Fail to list secrets")
		c.JSON(400, "Fail to list secrets")
		return
	}

	c.JSON(200, secrets)
}

func GetFile(mode string, name string) ([]byte, error) {
	suffix := ".encrypt"
	if mode == "encrypt" {
		suffix = ""
	}

	file := dataDir + name + suffix
	log.Info("file to lookup ", file)

	rawData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	vault := pkg.NewVault(name, crypter)

	var data []byte
	if mode == "encrypt" {
		data, err = vault.Encrypt(rawData)
		if err != nil {
			return nil, err
		}
	} else {
		data, err = vault.Decrypt(rawData)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
