package main

import (
	"errors"
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
		
		println(cmd.name + ": " + cmd.description)
	}
	return nil
}