package main

import "fmt"

func commandPodekex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You don't have any pokemon yet, go capture some")
		return nil
	}

	fmt.Println("Your Pokedex: ")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
