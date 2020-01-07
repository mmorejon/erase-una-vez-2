package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var (
	listenAddr string = "8000"
	character  string = "..."
)

func main() {
	router := mux.NewRouter()
	// register handlers
	router.HandleFunc("/healthz", healthzHandler)
	router.HandleFunc("/echo", echoHandler)
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

//
// Handlers
//

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, r, map[string]string{"status": "OK"})
	return
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// get hostname from os
	hostname, err := os.Hostname()
	if err != nil {
		log.Panicln("Error al obtener el hostname.")
	}
	// print message
	if len(os.Getenv("CHARACTER")) != 0 {
		character = os.Getenv("CHARACTER")
	}
	message := fmt.Sprintf("érase una vez %s", character)
	jsonResponse(w, r, map[string]string{"hostname": hostname, "message": message})
	return
}

//
// http utils
//

func jsonResponse(w http.ResponseWriter, r *http.Request, result interface{}) {
	body, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	w.Write(prettyJSON(body))
}

func prettyJSON(b []byte) []byte {
	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	return out.Bytes()
}
