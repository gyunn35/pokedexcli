package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}
	locationAreaName := args[0]
	response, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Printf("pokeman in %s:\n", locationAreaName)
	for _, v := range response.PokemonEncounters {
		fmt.Printf("	name: %s\n", v.Pokemon.Name)
	}

	return nil
}
