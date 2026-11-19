package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prp-e/go-for-rubyists/examples/internal/readinglist"
)

func main() {
	if err := run(); err != nil {
		slog.Error("reading list stopped with an error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return err
	}
	server := &http.Server{
		Handler:           readinglist.Handler(&readinglist.Store{}),
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	exited := make(chan error, 1)
	go func() { exited <- server.Serve(listener) }()
	slog.Info("reading list listening", "address", listener.Addr())

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
