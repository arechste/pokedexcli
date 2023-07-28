package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf(" - %s:%s\n", cmd.name, cmd.description)
	}
	return nil
}
