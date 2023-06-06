package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	for _, v := range cfg.caughtPokemon {
		fmt.Printf("	name: %s\n", v.Name)
	}
	return nil
}
