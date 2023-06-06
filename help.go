package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("welcome to the pokedex help menu!")
	fmt.Println("here are your available commands:")

	availableCommands := getCommands()

	for _, command := range availableCommands {
		fmt.Printf("	- %s: %s\n", command.name, command.desc)
	}

	return nil
}
