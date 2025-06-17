package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("inspect <pokemon_name>")
		return nil
	}

	pokemonName := args[0]
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pkType := range pokemon.Types {
			fmt.Printf("  - %s\n", pkType.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
