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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "              bye                universe            whatever",
			expected: []string{"bye", "universe", "whatever"},
		},
		{
			input:    "   hello",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Fatalf("for input %q, expected length %d but got %d", c.input, len(c.expected), len(actual))
		}

		for i := range actual {

			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("want %s, got %s", expectedWord, word)
			}
		}
	}
}
