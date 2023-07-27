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
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// cleanInput parses inputs from the REPL prompt.
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
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
	}
}

func prompt() {
	fmt.Print("Pokedex >")
}
