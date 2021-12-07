package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"training-microservice/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	echoHandler := handlers.NewEcho(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", echoHandler)

	server := &http.Server{
		Addr:         ":9999",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	sig := <-sigChannel
	logger.Println("Received terminate, gracefully shutting down", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
