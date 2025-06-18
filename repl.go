package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fernando8franco/pokedexcli/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, exist := getCommands()[words[0]]
		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := cmd.callback(cfg, args...)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the locations that can be visited within the games",
			callback:    commandMapNext,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous locations that can be visited within the games",
			callback:    commandMapPrevious,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Displays the pokemons that can be encountered in that location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Capture the specified pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspect the pokemon you have capture",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "See all the pokemons you have capture",
			callback:    commandPodekex,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
