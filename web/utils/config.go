package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	DBAddr string
	DBUser string
	DBPass string
	DBName string
	Addr string
}

func LoadConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	return &config, err
}