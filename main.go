package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func NewCli() map[string]cliCommand {
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

func commandHelp() error {
	//fmt.Printf("%v:%v\n")
	fmt.Println("help test")
	return nil
}

func commandExit() error {
	return nil
}

func prompt() {
	fmt.Print("Pokedex >")
}

func greeting() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
}

func main() {

	cli := NewCli()

	fmt.Println(cli)
	repl := bufio.NewScanner(os.Stdin)
	//fmt.Print("Pokedex >")
	prompt()
	for repl.Scan() {
		input := repl.Text()
		switch {
		case input == "help":
			greeting()
			fmt.Printf("%v:%v\n", cli[input].name, cli[input].description)
			fmt.Printf("%v:%v\n", cli["exit"].name, cli["exit"].description)
			fmt.Println("")

		case input == "exit":
			os.Exit(0)
		}
		prompt()
		//fmt.Println(repl.Text())
	}
	if err := repl.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
