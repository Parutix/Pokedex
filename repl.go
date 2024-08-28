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

		if text == "exit" {
			break
		}
	}
}