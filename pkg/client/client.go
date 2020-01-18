package client

import (
	"net/http"
)

// SetupClient is used to setup http client connection
func SetupClient() *http.Client {
	return &http.Client{}
}
