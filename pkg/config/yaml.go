package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Read(path string) (*Config, error) {
	var data []byte
	var cfg Config
	var err error

	if data, err = ioutil.ReadFile(path); err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	cfg.IsRelease = os.Getenv("GO_ENV") == "production"
	return &cfg, nil
}
