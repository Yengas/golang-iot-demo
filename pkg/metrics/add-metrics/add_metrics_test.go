package add_metrics_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"iot-demo/mocks"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	"iot-demo/pkg/metrics/alert"
	"iot-demo/pkg/metrics/ingestion"
	"testing"
	"time"
)

//go:generate mockgen -package mocks -destination ../../../mocks/add_metrics.go iot-demo/pkg/metrics/add-metrics Inserter,ConfigGetter

func TestAdd_Should_Fail_If_Insert_Fails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	deviceID, metricsToInsert := 52, []*ingestion.DecimalMetricValue(nil)
	expectedErr := errors.New("could not insert")

	inserter := mocks.NewMockInserter(ctrl)
	getter := mocks.NewMockConfigGetter(ctrl)

	addMetrics := add_metrics.NewService(inserter, getter)

	inserter.
		EXPECT().
		Insert(deviceID, metricsToInsert).
		Return(expectedErr).
		Times(1)

	getter.
		EXPECT().
		GetThreshold().
		Times(0)

	// Act
	res, err := addMetrics.Add(deviceID, metricsToInsert)

	// Assert
	assert.Empty(t, res)
	assert.Equal(t, err, expectedErr)
}

func TestAdd_Should_Succeed_And_Return_Ok(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	deviceID, metricsToInsert := 52, []*ingestion.DecimalMetricValue(nil)
	threshold := alert.Threshold{
		Min: 0,
		Max: 0,
	}

	inserter := mocks.NewMockInserter(ctrl)
	getter := mocks.NewMockConfigGetter(ctrl)

	addMetrics := add_metrics.NewService(inserter, getter)

	inserter.
		EXPECT().
		Insert(deviceID, metricsToInsert).
		Return(nil).
		Times(1)

	getter.
		EXPECT().
		GetThreshold().
		Return(threshold).
		Times(1)

	// Act
	res, err := addMetrics.Add(deviceID, metricsToInsert)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, res, add_metrics.Ok)
}


func TestAdd_Should_Alert_And_Return_Alert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Arrange
	deviceID, metricsToInsert := 52, []*ingestion.DecimalMetricValue{{
		Value: 5,
		Time:  ingestion.Time(time.Now()),
	}}
	threshold := alert.Threshold{
		Min: 1,
		Max: 10,
	}

	inserter := mocks.NewMockInserter(ctrl)
	getter := mocks.NewMockConfigGetter(ctrl)

	addMetrics := add_metrics.NewService(inserter, getter)

	inserter.
		EXPECT().
		Insert(deviceID, metricsToInsert).
		Return(nil).
		Times(1)

	getter.
		EXPECT().
		GetThreshold().
		Return(threshold).
		Times(1)

	// Act
	res, err := addMetrics.Add(deviceID, metricsToInsert)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, res, add_metrics.Alert)
}
