package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httClient: http.Client{
			Timeout: timeout,
		},
	}
}
