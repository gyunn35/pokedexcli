package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon provided")
	}
	pokemon := args[0]

	if v, ok := cfg.caughtPokemon[pokemon]; ok {
		fmt.Printf("	name: %s\n", v.Name)
		fmt.Printf("	height: %d\n", v.Height)
		fmt.Printf("	weight: %d\n", v.Weight)
		for _, stat := range v.Stats {
			fmt.Printf("		%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
	}

	return fmt.Errorf("need to caught pokemon %s before inspect", pokemon)
}
