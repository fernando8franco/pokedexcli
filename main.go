package main

import (
	"time"

	"github.com/fernando8franco/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(8*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
