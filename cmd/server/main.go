package main

import (
	"context"
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

func main() {
	cfgFile := "./resources/config.yaml"
	server, err := InitializeServer(cfgFile)
	if err != nil {
		log.Fatalf("could not read `%s` config file: %v\n", cfgFile, err)
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
