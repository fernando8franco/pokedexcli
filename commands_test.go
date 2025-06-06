package main

import (
	"testing"
)

func TestCommands(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "help",
			expected: "Displays a help message",
		},
		{
			input:    "exit",
			expected: "Exit the Pokedex",
		},
	}

	for _, c := range cases {
		if commands[c.input].description != c.expected {
			t.Errorf(
				"Actual description: %s --- Expected description: %s",
				commands[c.input].description,
				c.expected,
			)
		}
	}
}
