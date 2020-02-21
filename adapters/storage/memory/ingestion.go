package memory

import (
	"iot-demo/pkg/metrics/ingestion"
)

// IngestionDecimal is non threadsafe repository implementation for ingestion package
type IngestionDecimal struct {
	metrics map[int][]*ingestion.DecimalMetricValue
}

func (i *IngestionDecimal) Insert(deviceID int, metricsToInsert []*ingestion.DecimalMetricValue) error {
	for _, metric := range metricsToInsert {
		i.metrics[deviceID] = append(i.metrics[deviceID], metric)
	}
	return nil
}

func (i *IngestionDecimal) Query(deviceID int) ([]*ingestion.DecimalMetricValue, error) {
	if items, ok := i.metrics[deviceID]; ok {
		return items, nil
	}
	return nil, nil
}

func NewIngestion() *IngestionDecimal {
	metrics := make(map[int][]*ingestion.DecimalMetricValue)
	return &IngestionDecimal{metrics: metrics}
}
