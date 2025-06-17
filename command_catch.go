package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("catch <pokemon_name>")
		return nil
	}

	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	numRand := pokemon.BaseExperience
	if numRand <= 0 {
		numRand = 1
	}

	input := rand.Intn(numRand)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if input > 500 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	}

	return nil
}
