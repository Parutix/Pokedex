package main

import (
	"errors"
	"fmt"

	"github.com/Parutix/Pokedex/internal/pokeapi"
)

var commands map[string]command

type command struct {
	name        string
	description string
	function    func() error
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
			description: "Display a map of the region",
			function:		displayMap,
		},
	}
}

func exitPokedex() error {
	return errors.New("exit")
}

func listCommands() error {
	for _, cmd := range commands {

		if cmd.name == "" || cmd.description == "" {
			return errors.New("Error getting commands")
		}
		
		fmt.Println(cmd.name + ": " + cmd.description)
	}
	return nil
}

func displayMap() error {
	pokeAPIClient := pokeapi.NewClient()
	resp, err := pokeAPIClient.GetLocationAreas()
	if err != nil {
		return fmt.Errorf("Error getting location areas: %w", err)
	}
	for _, locationArea := range resp.Results {
		fmt.Println(locationArea.Name)
	}
	return nil
}
