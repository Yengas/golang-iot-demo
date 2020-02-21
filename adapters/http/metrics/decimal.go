package http_metrics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"iot-demo/pkg/auth"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	"iot-demo/pkg/metrics/ingestion"
	query_metrics "iot-demo/pkg/metrics/query-metrics"
	"strconv"
)

const (
	invalidDecimalMetricData = "invalid decimal metric data"
	invalidDeviceID          = "device id not valid"
	noMetricsFound           = "no metric found for given device and dates"
)

func decimalPostHandler(service *add_metrics.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj, exists := c.Get("auth_info")
		authInfo, ok := obj.(*auth.DeviceCredential)
		if !exists || !ok {
			c.String(401, "not authorized")
			return
		}

		var requestDTO []*ingestion.DecimalMetricValue
		if err := c.BindJSON(&requestDTO); err != nil {
			c.String(400, invalidDecimalMetricData)
			return
		}

		message, err := service.Add(authInfo.DeviceID, requestDTO)
		if err != nil {
			c.String(500, fmt.Errorf("insert error: %v", err).Error())
			return
		}

		c.JSON(200, message)
		return
	}
}

func decimalQueryHandler(service *query_metrics.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceIDSTR := c.Query("deviceID")
		deviceID, err := strconv.Atoi(deviceIDSTR)
		if err != nil {
			c.String(400, invalidDeviceID)
			return
		}

		metrics, err := service.Query(deviceID)
		if err != nil {
			c.String(500, err.Error())
			return
		}

		if metrics == nil {
			c.String(404, noMetricsFound)
			return
		}

		c.JSON(200, metrics)
		return
	}
}
