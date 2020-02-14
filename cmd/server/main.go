package main

import (
	"fmt"
	add_device "iot-demo/pkg/add-device"
	add_metrics "iot-demo/pkg/add-metrics"
	"iot-demo/pkg/ingestion"
	query_metrics "iot-demo/pkg/query-metrics"
	"iot-demo/pkg/registry"
	"iot-demo/pkg/storage/memory"
	"iot-demo/pkg/tokenizer"
	"time"
)

func main() {
	registryRepository := memory.NewRegistry()
	registryService := registry.NewService(registryRepository)

	ingestionRepository := memory.NewIngestion()
	ingestionService := ingestion.NewDecimalService(ingestionRepository)

	jwtConfig := tokenizer.Config{Secret: []byte("hello-world")}
	jwt := tokenizer.NewJWT(jwtConfig)

	deviceAdder := add_device.NewService(registryService, tokenizer.DeviceJWT(jwt))
	metricAdder := add_metrics.NewService(ingestionService)
	metricQuerier := query_metrics.NewService(ingestionService)

	device, token, err := deviceAdder.Register("asd", "aasd", time.Now())
	fmt.Printf("device: %v, token: %v, err: %v\n", device, token, err)

	got, ok := registryRepository.Get(device.ID)
	fmt.Printf("got device: %v, ok: %v\n", got, ok)

	metrics := []*ingestion.DecimalMetricValue{
		{0, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
	}

	err = metricAdder.Add(device.ID, metrics)
	fmt.Printf("insert err: %v\n", err)

	gotm, err := metricQuerier.Query(device.ID)
	fmt.Printf("queried metrics: %v, err: %v\n", gotm, err)
}
