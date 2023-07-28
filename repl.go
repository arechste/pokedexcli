package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// startRepl starts the cli REPL.
func startRepl(cfg *config) {
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
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		err := command.callback(cfg, args...)
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
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List Location Map next",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List Location Map previous",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "List pokemon for a location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon}",
			description: "Attempt to catch pokemon and add to your pokedex",
			callback:    callbackCatch,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

func prompt() {
	fmt.Print("Pokedex >")
}
