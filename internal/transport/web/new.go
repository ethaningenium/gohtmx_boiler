package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"time"
)

type Server struct {
	mux *http.ServeMux
}

func New() *Server {
	main := http.NewServeMux()
	handlers := NewHandlers()
	fsHandler := http.FileServer(http.Dir("./public"))
	main.Handle("/public/", http.StripPrefix("/public/", fsHandler))
	main.HandleFunc("/", handlers.Index)
	main.HandleFunc("/login", handlers.Login)
	main.HandleFunc("POST /todos", handlers.CreateTodo)
	main.HandleFunc("DELETE /todos/{id}", handlers.DeleteTodo)
	return &Server{mux: main}
}

func (s *Server) Serve() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Signal to close app
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	server := &http.Server{
		Addr:    ":4000",
		Handler: s.mux,
	}

	go func() {
		log.Printf("Starting server on port 4000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	select {
	case <-osSignal:
		log.Println("Shutdown signal received. Shutting down...")

		// Timeout to close handlers
		shutdownCtx, shutdownCancel := context.WithTimeout(ctx, 30*time.Second)
		defer shutdownCancel()

		// Stop the server
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server shutdown failed: %v", err)
		}
	}
}
