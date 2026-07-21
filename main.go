package main

import (
	"context"
	"errors"
	"lms/app"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	conn, err := app.NewConn(ctx)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	defer conn.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		Addr:              ":8000",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("listening on port %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case <-ctx.Done():
		log.Println("shutdown signal received")
	case err := <-errCh:
		if err != nil {
			log.Printf("server error: %v\n", err)
		}
		return
	}

	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("shutting down server")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed, forcing close: %v", err)
		if err := srv.Close(); err != nil {
			log.Printf("force close failed: %v", err)
		}
	} else {
		log.Println("graceful shutdown succeeded")
	}

	log.Println("shutdown complete")
}
