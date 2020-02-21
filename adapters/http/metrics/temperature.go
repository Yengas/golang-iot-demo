package http_metrics

import (
	"github.com/gin-gonic/gin"
	"iot-demo/pkg/metrics/ingestion"
)

// DeviceRegistry godoc
// @Summary insert temperature metric data for devices
// @Description inserts temperature metric data for the given device id
// @Security ApiKeyAuth
// @Tags metric
// @Accept json
// @Produce json
// @Param metrics body ingestion.DecimalMetricValueList true "metrics to insert"
// @Success 201 {string} string "inserted the temperature metrics"
// @Failure 400 {string} string "invalid request parameters"
// @Failure 401 {string} string "no device token supplied"
// @Failure 500 {string} string "unexpected error occurred"
// @Router /metric/temperature [post]
func noopTemperaturePostHandler(value ingestion.DecimalMetricValue) gin.HandlerFunc { return nil }

// DeviceRegistry godoc
// @Summary query temperature metrics of devices
// @Description given a device and a starting date, returns all temperature metrics
// @Tags metric
// @Accept json
// @Produce json
// @Param deviceID query string true "id of the device"
// @Success 200 {array} ingestion.DecimalMetricValue "metrics matching the criteria"
// @Failure 404 {string} string "no metrics found"
// @Failure 500 {string} string "unexpected error occurred"
// @Router /metric/temperature [get]
func noopTemperatureQueryHandler() gin.HandlerFunc { return nil }
