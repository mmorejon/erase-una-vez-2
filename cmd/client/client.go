package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mmorejon/erase-una-vez-2/pkg/client"
)

var (
	sleepTime string = "1s"
	endpoint  string = "http://localhost:8000/echo"
)

func main() {
	client := client.SetupClient()

	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime = os.Getenv("SLEEP_TIME")
	}

	if len(os.Getenv("ENDPOINT")) != 0 {
		endpoint = os.Getenv("ENDPOINT")
	}

	// Create a new request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Printf("Error al crear la petición a %s: %v\n", endpoint, err)
		os.Exit(1)
	}

	// Add headers from environment variables
	if headers := os.Getenv("HTTP_HEADERS"); headers != "" {
		// Headers format: "Key1:Value1,Key2:Value2"
		for _, header := range strings.Split(headers, ",") {
			parts := strings.Split(header, ":")
			if len(parts) == 2 {
				req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			}
		}
	}

	for {
		// Sleep time
		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)

		// Send the request
		r, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error al acceder a %s\n", endpoint)
			continue
		}

		// Read the response body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error al leer la respuesta de %s\n", endpoint)
			time.Sleep(sleepTimeDuration)
			continue
		}

		// Print the response body to stdout
		fmt.Printf("%s\n", body)
	}
}
