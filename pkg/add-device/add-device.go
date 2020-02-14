package add_device

import (
	"iot-demo/pkg/auth"
	"iot-demo/pkg/registry"
	"time"
)

type Registerer interface {
	Register(serialNumber string, firmwareVersion string, registrationDate time.Time) (*registry.Device, error)
}

type Tokenizer interface {
	Create(authInfo *auth.DeviceCredential) (auth.Token, error)
}

type Service struct {
	registerer Registerer
	tokenizer Tokenizer
}

func (s *Service) Register(serialNumber string, firmwareVersion string, registrationDate time.Time) (*registry.Device, auth.Token, error) {
	device, err := s.registerer.Register(serialNumber, firmwareVersion, registrationDate)
	if err != nil {
		return nil, "", err
	}

	authInfo := auth.DeviceCredential{DeviceID: device.ID}
	token, err := s.tokenizer.Create(&authInfo)
	if err != nil {
		return nil, "", err
	}

	return device, token, nil
}

func NewService(registerer Registerer, tokenizer Tokenizer) *Service {
	return &Service{
		registerer: registerer,
		tokenizer:  tokenizer,
	}
}
