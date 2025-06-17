package main

import (
	"time"

	"github.com/fernando8franco/pokedexcli/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 10*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	startRepl(cfg)
}
