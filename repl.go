package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned := cleanInput(strings.ToLower(scanner.Text()))

		if len(cleaned) == 0 {
			continue
		}

		fmt.Printf("Your command was: %s\n", cleaned[0])
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}
