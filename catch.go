package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	pokemonName := args[0]
	response, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	threshold := 50
	randNum := rand.Intn(response.BaseExperience)
	fmt.Printf("base experience: %d / randNum: %d\n", response.BaseExperience, randNum)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)

	cfg.caughtPokemon[pokemonName] = response

	return nil
}
