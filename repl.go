package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the Pokedex!")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := strings.TrimSpace(scanner.Text())

		if cmd, exists := commands[text]; exists {
			err := cmd.function()

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