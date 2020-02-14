package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

	return &cfg, nil
}
