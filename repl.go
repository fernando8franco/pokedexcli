package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		cmd, exist := commands[words[0]]
		if !exist {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	words := strings.Fields(lower)

	return words
}
