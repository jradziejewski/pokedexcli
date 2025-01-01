package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	initCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		cleaned := cleanInput(strings.ToLower(scanner.Text()))
		fmt.Println()
		if len(cleaned) == 0 {
			continue
		}

		fun := mapCommand(cleaned[0])
		fun()
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commandMap {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func mapCommand(commandName string) func() error {
	command, ok := commandMap[commandName]
	if ok {
		return command.callback
	} else {
		fmt.Println("unknown command")
	}

	return func() error { return nil }
}

var commandMap map[string]cliCommand

func initCommands() {
	commandMap = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
