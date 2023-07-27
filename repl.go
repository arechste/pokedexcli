package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl starts the cli REPL.
func startRepl() {
	for {
		reader := bufio.NewScanner(os.Stdin)
		//fmt.Print("Pokedex >")
		prompt()
		reader.Scan()
		cleaned := cleanInput(reader.Text())
		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

// cleanInput parses inputs from the REPL prompt.
func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	return words
}

// cliCommand struct data about supported commands and help.
type cliCommand struct {
	name        string
	description string
	callback    func() error
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
			description: "map forward",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "mapb backward",
			callback:    commandMapb,
		},
	}
}

func prompt() {
	fmt.Print("Pokedex >")
}
