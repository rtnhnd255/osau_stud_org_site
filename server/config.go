package server

import (
	"encoding/json"
	"os"

	"github.com/rtnhnd255/osau_stud_org_site.git/storage"
)

type Config struct {
	Addr  string `json:"port"`
	Store *storage.Config
}

func NewConfig(path string, store *storage.Config) (*Config, error) {
	var c Config
	cfgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(cfgFile)
	jsonParser.Decode(&c)
	c.Store = store
	return &c, nil
}
