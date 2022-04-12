package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	AppPort string `json:"appPort"`
}

func GetConfig(configFilePath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath.Clean(configFilePath))
	if err != nil {
		return nil, err
	}

	appConfig := &Config{}

	err = json.Unmarshal(configFile, &appConfig)
	if err != nil {
		return nil, err
	}

	return appConfig, nil
}
