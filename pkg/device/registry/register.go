package registry

import (
	"time"
)

type Repository interface {
	Register(serialNumber string, firmwareVersion string, registrationDate time.Time) (*Device, error)
}

type Service struct {
	repository Repository
}

func (s *Service) Register(serialNumber string, firmwareVersion string, registrationDate time.Time) (*Device, error) {
	return s.repository.Register(serialNumber, firmwareVersion, registrationDate)
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
