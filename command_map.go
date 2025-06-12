package main

import (
	"fmt"
)

func commandMapNext(cfg *config) error {
	url := cfg.Next

	locations, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapPrevious(cfg *config) error {
	url := cfg.Previous

	locations, err := cfg.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
