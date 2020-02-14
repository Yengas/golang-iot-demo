package ingestion

type DecimalMetricValue struct {
	Value float64 `json:"value"`
	// Epoch timestamp in seconds
	Time Time `json:"time" swaggertype:"primitive,integer" example:"1578859629"`
}

type DecimalRepository interface {
	Insert(deviceID int, metricsToInsert []*DecimalMetricValue) error
	Query(deviceID int) ([]*DecimalMetricValue, error)
}

type DecimalService struct {
	repository DecimalRepository
}

// insert metrics into the backing repository, notify all listeners
func (ss *DecimalService) Insert(deviceID int, metricsToInsert []*DecimalMetricValue) error {
	err := ss.repository.Insert(deviceID, metricsToInsert)
	if err != nil {
		return err
	}

	return nil
}

func (ss *DecimalService) Query(deviceID int) ([]*DecimalMetricValue, error) {
	return ss.repository.Query(deviceID)
}

func NewDecimalService(repository DecimalRepository) *DecimalService {
	return &DecimalService{repository: repository}
}
