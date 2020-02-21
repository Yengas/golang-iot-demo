package http_server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	http_device "iot-demo/adapters/http/device"
	"iot-demo/adapters/http/docs"
	http_health "iot-demo/adapters/http/health"
	http_metrics "iot-demo/adapters/http/metrics"
	add_device "iot-demo/pkg/device/add-device"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	query_metrics "iot-demo/pkg/metrics/query-metrics"
	"log"
	"net/http"
)

//go:generate swag init --parseDependency -g server.go
// @title IOT Demo
// @version 1.0
// @description Devices and Metrics API

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @contact.name Yengas
// @contact.email yigitcan.ucum@trendyol.com

// @Schemes http

type httpHandleServer interface {
	ListenAndServe() error
	Shutdown(ctx context.Context) error
}

type Config struct {
	Host                  string
	DocumentationHost     string
	DocumentationBasePath string
	Port                  int
	IsRelease             bool
}

type HTTPServer struct {
	config Config
	server httpHandleServer
}

type InstrumentationHandler gin.HandlerFunc

func NewHandlers(
	config Config,
	tokenParser DeviceTokenParser,
	addDevice *add_device.Service,
	addMetrics *add_metrics.Service,
	queryMetrics *query_metrics.Service,
) http.Handler {
	if config.IsRelease {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := createGinEngine()

	engine.Use(deviceAuthParserHandler(tokenParser))

	docs.SwaggerInfo.Host = config.DocumentationHost
	docs.SwaggerInfo.BasePath = config.DocumentationBasePath

	engine.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("", func(c *gin.Context) {
		c.Redirect(302, "/swagger-ui/index.html")
		c.Abort()
	})

	http_health.Handlers{}.Register(engine)
	http_device.Handlers{Service: addDevice}.Register(engine)
	http_metrics.Handlers{AddMetricsService: addMetrics, QueryMetricsService: queryMetrics}.Register(engine)

	return engine
}

func createGinEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	return engine
}

func (httpServer *HTTPServer) Start() error {
	cfg := httpServer.config
	log.Printf("server is starting on %v:%v\n", cfg.Host, cfg.Port)
	return httpServer.server.ListenAndServe()
}

func (httpServer *HTTPServer) Stop(ctx context.Context) error {
	return httpServer.server.Shutdown(ctx)
}

func createServer(cfg Config, handler http.Handler) *http.Server {
	listenAddress := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	return &http.Server{Addr: listenAddress, Handler: handler}
}

func NewServer(
	config Config,
	handler http.Handler,
) *HTTPServer {
	return &HTTPServer{config: config, server: createServer(config, handler)}
}
