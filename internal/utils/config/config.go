package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ValidApiKeys map[string]string `json:"validApiKeys"`
}

func GetConfig(path string) (*Config, error)  {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Could not read file from path %s. %v", path, err)
	}
	
	config, err := ParseConfig(data)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func ParseConfig(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("ERROR: could not parse config file. %v", err)
	}

	return &config, nil
}