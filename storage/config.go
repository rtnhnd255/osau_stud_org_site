package storage

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBURL string `json:"db_url"`
}

func NewConfig(path string) (*Config, error) {
	var c Config
	cfgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(cfgFile)
	jsonParser.Decode(&c)
	return &c, nil
}
