package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	podName := os.Getenv("POD_NAME")

	mux := http.NewServeMux()
	mux.HandleFunc("/__gtg", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	})

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		msg := fmt.Sprintf("POD_NAME: %s\n", podName)
		rw.Write([]byte(msg))

	})

	srv := &http.Server{
		ReadTimeout:  25 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
		Addr:         ":8080",
		Handler:      mux,
	}

	idleConnsClosed := make(chan struct{}, 1)
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, syscall.SIGTERM) // k8s sends SIGTERM
		<-sigint

		srv.SetKeepAlivesEnabled(false)

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Unable to start HTTP server")
	}

	<-idleConnsClosed
}
