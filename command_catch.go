package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	// not found error (404) empty pokemon  err not catched
	fmt.Println(pokemon)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		fmt.Printf("failed to catch %s\n", pokemon.Name)
		return nil

	}
	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon

	return nil

}
