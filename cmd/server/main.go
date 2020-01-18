package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/mmorejon/erase-una-vez-2/pkg/server"
)

var (
	listenAddr string = "8000"
)

func main() {
	router := mux.NewRouter()
	// register handlers
	router.HandleFunc("/healthz", server.HealthzHandler).Methods("GET")
	router.HandleFunc("/echo", server.EchoHandler).Methods("GET")
	// h2c configuration
	handler := h2c.NewHandler(router, &http2.Server{})

	// server configuration
	server := &http.Server{
		Handler:      handler,
		Addr:         ":" + listenAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// start server
	log.Println("Servidor iniciado")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}
