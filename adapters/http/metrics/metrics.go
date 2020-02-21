package http_metrics

import (
	"github.com/gin-gonic/gin"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	query_metrics "iot-demo/pkg/metrics/query-metrics"
)

type Handlers struct {
	AddMetricsService   *add_metrics.Service
	QueryMetricsService *query_metrics.Service
}

func (drh Handlers) Register(engine *gin.Engine) {
	engine.POST("/metric/temperature", decimalPostHandler(drh.AddMetricsService))
	engine.GET("/metric/temperature", decimalQueryHandler(drh.QueryMetricsService))
}
