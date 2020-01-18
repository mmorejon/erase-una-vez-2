package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mmorejon/erase-una-vez-2/pkg/utils"
)

var (
	character string = "..."
)

// HealthzHandler is used to response the service health
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, r, map[string]string{"status": "OK"})
	return
}

// EchoHandler is used to print a message
func EchoHandler(w http.ResponseWriter, r *http.Request) {
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
	utils.JSONResponse(w, r, map[string]string{"hostname": hostname, "message": message})
	return
}
