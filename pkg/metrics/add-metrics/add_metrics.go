package add_metrics

import (
	"iot-demo/pkg/metrics/alert"
	"iot-demo/pkg/metrics/ingestion"
)

type ResponseToken string
const (
	Alert ResponseToken = "ALERT"
	Ok ResponseToken = "OK"
)

type Inserter interface {
	Insert(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) error
}

type ConfigGetter interface {
	GetThreshold() alert.Threshold
}

type Service struct {
	inserter Inserter
	config   ConfigGetter
}

func (s *Service) Add(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) (ResponseToken, error) {
	err := s.inserter.Insert(deviceID, metricsToInsert)
	if err != nil {
		return "", err
	}

	threshold := s.config.GetThreshold()
	return getMessage(threshold, metricsToInsert), nil
}

func getMessage(threshold alert.Threshold, metrics []*ingestion.DecimalMetricValue) ResponseToken {
	for _, metric := range metrics {
		if threshold.DoesMatch(metric.Value) {
			return Alert
		}
	}
	return Ok
}

func NewService(inserter Inserter, config ConfigGetter) *Service {
	return &Service{inserter: inserter, config: config}
}
