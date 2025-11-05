package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alechekz/online-car-auction/services/vehicle/internal/server"
)

// main initializes and starts the Vehicle Service HTTP server
func main() {

	// Prepare server
	cfg := server.NewConfig(":" + os.Getenv("VEHICLE_PORT"))
	srv := server.NewServer(cfg)

	// Graceful shutdown handling
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a separate goroutine
	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	<-stop
	if err := srv.Stop(); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}

}
