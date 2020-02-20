package main

import (
	"context"
	"iot-demo/pkg/config"
	http_server "iot-demo/pkg/http"
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

func getConfigProfile() string {
	env := os.Getenv("PROFILE")
	if env != "" {
		return env
	}
	return "dev"
}

func main() {
	cfgSelection := config.Selection{
		StaticConfigurationFilePath:  "./resources/config.yaml",
		DynamicConfigurationFilePath: "./resources/config_dynamic.yaml",
		Profile:                      getConfigProfile(),
	}
	server, err := InitializeServer(cfgSelection)
	if err != nil {
		log.Fatalf("could not read `%v` config file: %v\n", cfgSelection, err)
	}
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
