package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mmorejon/erase-una-vez-2/pkg/server"
)

var (
	listenAddr string = ":8000"
)

func main() {
	router := mux.NewRouter()
	// register handlers
	router.HandleFunc("/healthz", server.HealthzHandler).Methods("GET")
	router.HandleFunc("/echo", server.EchoHandler).Methods("GET")

	// start server
	log.Println("Servidor iniciado")
	err := http.ListenAndServe(listenAddr, router)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}
}
