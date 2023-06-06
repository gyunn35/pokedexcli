package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	desc     string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:     "help",
			desc:     "display the help menu",
			callback: commandHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "terminate the pokedex",
			callback: commandExit,
		},
		"map": {
			name:     "map",
			desc:     "display list of location areas",
			callback: commandMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "display list of location areas that has seen before",
			callback: commandMapb,
		},
		"explore": {
			name:     "explore",
			desc:     "lists the pokemon in a location area",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch",
			desc:     "try to catch a pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name:     "inspect",
			desc:     "try to inspect the caught pokemon",
			callback: commandInspect,
		},
		"pokedex": {
			name:     "pokedex",
			desc:     "view all the pokemon in your pokedex",
			callback: commandPokedex,
		},
	}
}

func getInput(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCommands()

	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		input := scanner.Text()

		cleaned := cleanInput(input)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		var args []string
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	lowered := strings.ToLower(input)
	words := strings.Fields(lowered) // split input by white space
	return words
}
