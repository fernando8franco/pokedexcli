package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello   world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "THIS is A TEST",
			expected: []string{"this", "is", "a", "test"},
		},
		{
			input:    "Gastly   Haunter   Gengar",
			expected: []string{"gastly", "haunter", "gengar"},
		},
		{
			input:    " ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(
				"Actual length: %v --- Expected length: %v",
				len(actual),
				len(c.expected),
			)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"Actual word: %s --- Expected word: %s",
					word,
					expectedWord,
				)
			}
		}
	}
}
