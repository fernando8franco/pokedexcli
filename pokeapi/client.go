package pokeapi

import (
	"net/http"
	"time"

	"github.com/fernando8franco/pokedexcli/pokecache"
)

type Client struct {
	httClient http.Client
	cache     pokecache.Cache
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		httClient: http.Client{
			Timeout: timeout,
		},
		cache: *pokecache.NewCache(cacheInterval),
	}
}
