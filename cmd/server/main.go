package main

import (
	"context"
	"fmt"
	add_device "iot-demo/pkg/device/add-device"
	"iot-demo/pkg/device/registry"
	http_server "iot-demo/pkg/http"
	add_metrics "iot-demo/pkg/metrics/add-metrics"
	"iot-demo/pkg/metrics/ingestion"
	query_metrics "iot-demo/pkg/metrics/query-metrics"
	"iot-demo/pkg/storage/memory"
	"iot-demo/pkg/tokenizer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startServer(server *http_server.HTTPServer) {
	if err := server.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not start the application: %v", err.Error())
	}
}

func createGracefulShutdownChannel() chan os.Signal {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM)
	signal.Notify(gracefulShutdown, syscall.SIGINT)
	return gracefulShutdown
}

func main() {
	registryRepository := memory.NewRegistry()
	registryService := registry.NewService(registryRepository)

	ingestionRepository := memory.NewIngestion()
	ingestionService := ingestion.NewDecimalService(ingestionRepository)

	jwtConfig := tokenizer.Config{Secret: []byte("hello-world")}
	jwt := tokenizer.NewJWT(jwtConfig)

	addDevice := add_device.NewService(registryService, tokenizer.DeviceJWT(jwt))
	addMetrics := add_metrics.NewService(ingestionService)
	queryMetrics := query_metrics.NewService(ingestionService)

	device, token, err := addDevice.Register("asd", "aasd", time.Now())
	fmt.Printf("device: %v, token: %v, err: %v\n", device, token, err)

	got, ok := registryRepository.Get(device.ID)
	fmt.Printf("got device: %v, ok: %v\n", got, ok)

	metrics := []*ingestion.DecimalMetricValue{
		{0, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
		{1, ingestion.Time(time.Now())},
	}

	err = addMetrics.Add(device.ID, metrics)
	fmt.Printf("insert err: %v\n", err)

	gotm, err := queryMetrics.Query(device.ID)
	fmt.Printf("queried metrics: %v, err: %v\n", gotm, err)

	serverConfig := http_server.Config{
		Host:                  "localhost",
		DocumentationHost:     "",
		DocumentationBasePath: "/",
		Port:                  8080,
		IsRelease:             false,
	}
	handlers := http_server.NewHandlers(serverConfig, tokenizer.DeviceJWT(jwt), addDevice, addMetrics, queryMetrics)
	server := http_server.NewServer(serverConfig, handlers)
	gracefulShutdown := createGracefulShutdownChannel()
	// start the server and graceful shutdown
	go startServer(server)

	sig := <-gracefulShutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("caught sig: %+v, shutting down the application\n", sig)
	if err = server.Stop(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the application: %v", err.Error())
	}
	log.Println("shutdown application gracefully.")
}
