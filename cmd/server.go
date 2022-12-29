package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/tmp-friends/victo-api/app/config"
)

func main() {
	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// TODO: .env

	addr := config.LoadConfig().HTTPInfo.Addr

	// TODO: DI container settings

	// TODO: root settings
	router := config.InitRouter()

	// Create app
	s := http.Server{
		Addr:    addr,
		Handler: router,
	}

	// Graceful shutdown
	go func() {
		// @see https://echo.labstack.com/guide/http_server/
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
		log.Println("Server is running! addr: ", addr)
	}()

	// Listen for the interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown
	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 second to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
