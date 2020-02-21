package add_metrics

import (
	"iot-demo/pkg/metrics/ingestion"
)

type Inserter interface {
	Insert(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) error
}

type ConfigGetter interface {
	GetMessage() string
}

type Service struct {
	inserter Inserter
	getter   ConfigGetter
}

func (s *Service) Add(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) (string, error) {
	err := s.inserter.Insert(deviceID, metricsToInsert)
	if err != nil {
		return "", err
	}

	return s.getter.GetMessage(), nil
}

func NewService(inserter Inserter, getter ConfigGetter) *Service {
	return &Service{inserter: inserter, getter: getter}
}
