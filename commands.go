package main

import (
	"errors"
	"fmt"
)

var commands map[string]command

type command struct {
	name        string
	description string
	function    func(*config) error
}

func initCommands() {
	commands = map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			function:    exitPokedex,
		},
		"help": {
			name:				"help",
			description: "List all available commands",
			function:		listCommands,
		},
		"map": {
			name:				"map",
			description: "Display a map of the region. If the user writes map again, it will display the next page of the map",
			function:		displayMap,
		},
		"mapb": {
			name:				"mapb",
			description: "Display the previous page of the map, if there is no previous page, it will display the first page",
			function:		displayMapb,
		},
	}
}

func exitPokedex(cfg *config) error {
	return errors.New("exit")
}

func listCommands(cfg *config) error {
	for _, cmd := range commands {

		if cmd.name == "" || cmd.description == "" {
			return errors.New("Error getting commands")
		}
		
		fmt.Println(cmd.name + ": " + cmd.description)
	}
	return nil
}

func displayMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w", err)
	}
	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}
	cfg.nextLocationAreasURL = resp.Next
	cfg.previousLocationAreasURL = resp.Previous
	return nil
}

func displayMapb(cfg *config) error {
	resp, err := cfg.pokeapiClient.GetLocationAreas(cfg.previousLocationAreasURL)
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w", err)
	}
	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}
	cfg.nextLocationAreasURL = resp.Next
	cfg.previousLocationAreasURL = resp.Previous
	return nil
}
