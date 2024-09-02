package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var commandsWithArgs = map[string]bool{
	"explore": true,
}

func splitCommand(input string) (string, string) {
	parts := strings.SplitN(input, " ", 2)
	if len(parts) < 2 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		cmdName, rawArgs := splitCommand(text)
		if _, exists := commandsWithArgs[cmdName]; exists && rawArgs == "" {
			fmt.Println("Command requires arguments.")
			continue
		}
		
		args := strings.Split(rawArgs, " ")
		if cmd, exists := commands[cmdName]; exists {
			err := cmd.function(cfg, args...)

			if err != nil {
				if(err.Error() == "exit") {

					fmt.Println("Exiting Pokedex...")
					break

				} else {

					fmt.Println("Caught Error: ", err)
					
				}
			}
		} else {
			fmt.Println("Command not found.")
		} 
	}
}