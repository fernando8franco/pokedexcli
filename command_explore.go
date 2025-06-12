package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	if len(cfg.Params) < 1 {
		fmt.Println("explore <area_name>")
		return nil
	}

	pokemons, err := cfg.pokeapiClient.ListPokemonsInLocation(cfg.Params[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", cfg.Params[0])
	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
