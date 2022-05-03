package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/go-pkgz/lgr"
	"github.com/the-NZA/tuff-stuff/backend/internal/config"
	"github.com/the-NZA/tuff-stuff/backend/internal/store"
	"github.com/the-NZA/tuff-stuff/backend/internal/tuff"
)

var (
	configPath string
	doneChan   = make(chan struct{}, 1)
	timeout    = time.Second * 15
)

func init() {
	flag.StringVar(&configPath, "config", "etc/dev.json", "Path to config file")
}

func main() {
	if err := run(); err != http.ErrServerClosed {
		log.Fatalf("[ERROR] %v", err)
	}

	<-doneChan

	log.Printf("[INFO] Application was shutdowned\n")
}

// run creates and starts application.
func run() error {
	flag.Parse()

	// Load application config
	conf, err := config.NewConfigFromFile(configPath)
	if err != nil {
		return err
	}

	// Create store
	s, err := store.NewStore(conf)
	if err != nil {
		return err
	}

	// Create application
	app := tuff.NewApp(conf, s)

	// Start shutdown goroutine
	go shutdown(app, doneChan)

	// Start application
	return app.Start()
}

// shutdown stops application.
func shutdown(app tuff.Apper, done chan struct{}) {
	// Register signal listener
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan

	log.Printf("[INFO] Application is shutting down...\n")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Try to graceful shutdown
	if err := app.Stop(ctx); err != nil && err != http.ErrServerClosed {
		log.Printf("[ERROR] Application can't be gracefully shutdowned: %v\n", err)
	}

	close(done)
}
