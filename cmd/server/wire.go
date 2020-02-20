//+build wireinject

package main

import (
	"github.com/google/wire"
	"iot-demo/pkg/config"
	add_device "iot-demo/pkg/device/add-device"
	"iot-demo/pkg/device/registry"
	http_server "iot-demo/pkg/http"
	"iot-demo/pkg/jwt"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	"iot-demo/pkg/metrics/ingestion"
	query_metrics "iot-demo/pkg/metrics/query-metrics"
	"iot-demo/pkg/storage/memory"
)

func NewJWTConfig(cfg config.Static) jwt.Config {
	return jwt.Config{Secret: []byte(cfg.Auth.Secret)}
}

func NewDeviceJWT(authJWT jwt.AuthJWT) jwt.DeviceJWT {
	return jwt.DeviceJWT(authJWT)
}

func NewHTTPServerConfig(cfg config.Static) http_server.Config {
	return http_server.Config{
		Host:                  cfg.Server.Host,
		DocumentationHost:     cfg.Swagger.DocumentationHost,
		DocumentationBasePath: cfg.Swagger.DocumentationBasePath,
		Port:                  cfg.Server.Port,
		IsRelease:             cfg.IsRelease,
	}
}

var (
	configSet = wire.NewSet(
		config.ReadStatic,
		config.NewDynamicGetter,
		wire.Bind(new(add_metrics.ConfigGetter), new(*config.DynamicGetter)),
	)
	jwtSet = wire.NewSet(
		NewJWTConfig,
		jwt.NewJWT,
		NewDeviceJWT,
		wire.Bind(new(http_server.DeviceTokenParser), new(jwt.DeviceJWT)),
		wire.Bind(new(add_device.Tokenizer), new(jwt.DeviceJWT)),
	)
	storageSet = wire.NewSet(
		memory.NewIngestion,
		memory.NewRegistry,
		wire.Bind(new(registry.Repository), new(*memory.Registry)),
		wire.Bind(new(ingestion.DecimalRepository), new(*memory.IngestionDecimal)),
	)
	httpSet = wire.NewSet(
		NewHTTPServerConfig,
		http_server.NewHandlers,
		http_server.NewServer,
	)
	deviceSet = wire.NewSet(
		registry.NewService,
		wire.Bind(new(add_device.Registerer), new(*registry.Service)),
	)
	ingestionSet = wire.NewSet(
		ingestion.NewDecimalService,
		wire.Bind(new(add_metrics.Inserter), new(*ingestion.DecimalService)),
		wire.Bind(new(query_metrics.Querier), new(*ingestion.DecimalService)),
	)
	addDeviceSet    = wire.NewSet(add_device.NewService)
	addMetricsSet   = wire.NewSet(add_metrics.NewService)
	queryMetricsSet = wire.NewSet(query_metrics.NewService)
)

func InitializeServer(cs config.Selection) (*http_server.HTTPServer, error) {
	wire.Build(
		configSet,
		jwtSet,
		deviceSet,
		ingestionSet,
		storageSet,
		addDeviceSet,
		addMetricsSet,
		queryMetricsSet,
		httpSet,
	)
	return nil, nil
}
