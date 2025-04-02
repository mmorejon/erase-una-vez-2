package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mmorejon/erase-una-vez-2/pkg/client"
)

var (
	sleepTime string = "1s"
	endpoint string = "http://localhost:8000/echo"
)

func main() {
	client := client.SetupClient()

	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime = os.Getenv("SLEEP_TIME")
	}

	if len(os.Getenv("ENDPOINT")) != 0 {
		endpoint = os.Getenv("ENDPOINT")
	}

	for {
		// Sleep time
		sleepTimeDuration, _ := time.ParseDuration(sleepTime)
		time.Sleep(sleepTimeDuration)

		// Get the endpoint
		r, err := client.Get(endpoint)
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
