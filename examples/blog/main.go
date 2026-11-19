package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	startup, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	db, err := openStore(startup, "blog.db")
	cancel()
	if err != nil {
		return err
	}
	defer db.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:              "127.0.0.1:8081",
		Handler:           routes(db),
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}
	exited := make(chan error, 1)
	go func() { exited <- server.ListenAndServe() }()
	log.Printf("Open http://%s", server.Addr)

	select {
	case err := <-exited:
		return err
	case <-ctx.Done():
		stop()
	}
	shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdown); err != nil {
		return errors.Join(err, server.Close())
	}
	if err := <-exited; !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
