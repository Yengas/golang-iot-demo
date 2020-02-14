package query_metrics

import (
	"iot-demo/pkg/ingestion"
)

type Querier interface {
	Query(deviceID int) ([]*ingestion.DecimalMetricValue, error)
}

type Service struct {
	querier Querier
}

func (s *Service) Query(deviceID int) ([]*ingestion.DecimalMetricValue, error) {
	return s.querier.Query(deviceID)
}

func NewService(querier Querier) *Service {
	return &Service{querier: querier}
}

