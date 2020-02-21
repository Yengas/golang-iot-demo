package config

import (
	"errors"
	"github.com/fsnotify/fsnotify"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"iot-demo/pkg/metrics/alert"
	"sync"
)

type DynamicGetter struct {
	viperInstance *viper.Viper
	mutex         *sync.RWMutex
	config        *Dynamic
}


func (dg *DynamicGetter) GetThreshold() alert.Threshold {
	return dg.getCurrent().Threshold
}

func NewDynamicGetter(s Selection) (*DynamicGetter, error) {
	viperInstance := viper.New()
	viperInstance.SetConfigFile(s.DynamicConfigurationFilePath)
	viperInstance.SetTypeByDefaultValue(true)
	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, errors.New("config could not be found")
	}
	viperInstance.WatchConfig()
	dg := &DynamicGetter{viperInstance, &sync.RWMutex{}, nil}

	viperInstance.OnConfigChange(func(e fsnotify.Event) {
		if err := dg.Update(); err != nil {
			log.Error("an error happened when updating the config: " + err.Error())
		}
	})

	return dg, dg.Update()
}

func (dg *DynamicGetter) Update() error {
	var cfg Dynamic
	err := dg.viperInstance.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	dg.mutex.Lock()
	defer dg.mutex.Unlock()
	dg.config = &cfg
	return nil
}

func (dg *DynamicGetter) getCurrent() Dynamic {
	dg.mutex.RLock()
	defer dg.mutex.RUnlock()
	return *dg.config
}
