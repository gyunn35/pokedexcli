package main

import (
	"fmt"
)

func commandMap(cfg *config, args ...string) error {
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("location areas:")
	for _, v := range response.Results {
		fmt.Printf("	name: %s (%s)\n", v.Name, v.URL)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("there is no record of previously collating the list; please run map command first")
	}
	response, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("location areas:")
	for _, v := range response.Results {
		fmt.Printf("	name: %s (%s)\n", v.Name, v.URL)
	}

	cfg.nextLocationAreaURL = response.Next
	cfg.previousLocationAreaURL = response.Previous

	return nil
}
