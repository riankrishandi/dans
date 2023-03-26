package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/riankrishandi/dans/controller"
	"github.com/riankrishandi/dans/job"
	"github.com/riankrishandi/dans/middleware"
	"github.com/riankrishandi/dans/render"
	"github.com/riankrishandi/dans/repo"
	"github.com/riankrishandi/dans/route"
	"github.com/riankrishandi/dans/server"
)

func main() {
	// Init repo.
	r, err := repo.NewRepo()
	if err != nil {
		log.Fatalf("failed to init repo: %s", err.Error())
	}

	// Init renderer.
	renderer := render.New()

	// Init job.
	job := job.New()

	// Init controller.
	controller := controller.New(r, job, renderer)

	// Init middleware.
	m := middleware.New(renderer)

	// Init route.
	mux, err := route.NewAPI(controller, m)
	if err != nil {
		log.Fatalf("failed to init api: %s", err.Error())
	}

	// Initialize server.
	srv := server.New(mux)

	// Start HTTP server.
	go func() {
		srv.ServeHTTPHandler()
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		shutdownRelease()
	}()

	srv.ShutdownHTTPHandler(shutdownCtx)
	log.Println("[main] graceful shutdown complete")
}
