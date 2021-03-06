package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	AppPort         string `json:"appPort"`
	AWSID           string `json:"awsID"`
	AWSSecret       string `json:"awsSecret"`
	AWSToken        string `json:"awsToken"`
	AWSHost         string `json:"awsHost"`
	AWSRegion       string `json:"awsRegion"`
	MoviesTableName string `json:"moviesTableName"`
}

func GetConfig(configFilePath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath.Clean(configFilePath))
	if err != nil {
		return nil, fmt.Errorf("config.GetConfig(): %w", err)
	}

	appConfig := &Config{}

	err = json.Unmarshal(configFile, &appConfig)
	if err != nil {
		return nil, fmt.Errorf("config.GetConfig(): %w", err)
	}

	return appConfig, nil
}
