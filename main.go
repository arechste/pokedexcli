package main

import (
	"time"

	"github.com/arechste/pokedexcli/internal/pokeapi"
)

// timeout is time for apiClient timeout.
const timeout = 5 * time.Second

// interval is time for cache flush interval
const interval = 5 * time.Minute

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(timeout, interval),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
