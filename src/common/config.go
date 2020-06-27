package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"path"
	"path/filepath"
)

type Config struct {
	ConnString string
}

func NewConfig() *Config {
	c := &Config{}
	configure(c)
	return c
}

func configure(config *Config) {
	config.Load("./configs")
	config.Load("./secrets")
}

// This config loader can only accept json file
func (c *Config) Load(mountedConfigDir string) error {
	var buffer bytes.Buffer

	files, err := ioutil.ReadDir(mountedConfigDir)

	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fullFilePath := path.Join(mountedConfigDir, file.Name())
			if loadConfig(fullFilePath, c) != nil {
				buffer.WriteString("Cannot load the secret file\n")
			}
		}
	}

	if buffer.Len() > 0 {
		return errors.New(buffer.String())
	}

	return nil
}

func loadConfig(file string, config *Config) error {
	data, _ := ioutil.ReadFile(file)
	return json.Unmarshal(data, &config)
}
