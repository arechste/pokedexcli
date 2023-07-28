package pokeapi

import (
	"net/http"
	"time"

	"github.com/arechste/pokedexcli/internal/pokecache"
)

// baseUrl of pokemon API
const baseUrl = "https://pokeapi.co/api/v2"

// Client HTTP/API.
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
