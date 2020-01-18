package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mmorejon/erase-una-vez-2/pkg/client"
)

var (
	sleepTime int    = 1
	serverURL string = "http://localhost:8000"
)

func main() {
	client := client.SetupClient()

	if len(os.Getenv("SLEEP_TIME")) != 0 {
		sleepTime, _ = strconv.Atoi(os.Getenv("SLEEP_TIME"))
	}

	if len(os.Getenv("SERVER_URL")) != 0 {
		serverURL = os.Getenv("SERVER_URL")
	}

	endpoint := os.Getenv("ENDPOINT")

	for {
		accessPoint := fmt.Sprint(serverURL, endpoint)
		r, err := client.Get(accessPoint)
		if err != nil {
			log.Fatal(err)
		}

		// Read the response body
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Print the response body to stdout
		fmt.Printf("%s\n", body)

		// Sleep
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}
}
