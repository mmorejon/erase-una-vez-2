package client

import (
	"net"
	"net/http"
	"time"
)

// SetupClient is used to setup http client connection
func SetupClient() *http.Client {

	return &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   5 * time.Second,
				KeepAlive: 5 * time.Second,
			}).DialContext,
			DisableKeepAlives:     true,
			TLSHandshakeTimeout: 5 * time.Second,
			ResponseHeaderTimeout: 5 * time.Second,
		},
	}
}
