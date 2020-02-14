package add_metrics

import (
	"iot-demo/pkg/ingestion"
)

type Inserter interface {
	Insert(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) error
}

type Service struct {
	inserter Inserter
}

func (s *Service) Add(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) error {
	return s.inserter.Insert(deviceID, metricsToInsert)
}

func NewService(inserter Inserter) *Service {
	return &Service{inserter: inserter}
}

