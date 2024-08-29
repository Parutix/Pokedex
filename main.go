package main

import "github.com/Parutix/Pokedex/internal/pokeapi"

type config struct {
	pokeapiClient 					 pokeapi.Client
	nextLocationAreasURL 		 *string
	previousLocationAreasURL *string
}

func main() {
	cfg := config {
		pokeapiClient: pokeapi.NewClient(),
	}
	initCommands()
	startREPL(&cfg)
}