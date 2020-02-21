package config

import (
	"errors"
	"github.com/spf13/viper"
	"os"
)

func ReadStatic(s Selection) (Static, error) {
	viperInstance := viper.New()
	viperInstance.SetConfigFile(s.StaticConfigurationFilePath)

	err := viperInstance.ReadInConfig()
	if err != nil {
		return Static{}, err
	}

	cfg, err := createSubConfiguration(viperInstance, s.Profile)
	if err != nil {
		return Static{}, err
	}

	cfg.IsRelease = os.Getenv("GO_ENV") == "production"
	return cfg, err
}

func createSubConfiguration(viperInstance *viper.Viper, sub string) (Static, error) {
	configuration := Static{}
	subEnvironment := viperInstance.Sub(sub)
	if subEnvironment == nil {
		return Static{}, errors.New("no config found for the given environment")
	}

	defaultEnvironment := viperInstance.Sub("default")
	if defaultEnvironment == nil {
		return Static{}, errors.New("no config found for the default environment")
	}

	defaultEnvironment.MergeConfigMap(subEnvironment.AllSettings())

	defaultEnvironment.BindEnv("server.port", "PORT")

	err := defaultEnvironment.Unmarshal(&configuration)
	if err != nil {
		return Static{}, nil
	}
	return configuration, nil
}
