package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

var commands map[string]command

type command struct {
	name        string
	description string
	function    func(*config, ...string) error
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
		"explore": {
			name:				"explore",
			description: "Explore the region and find new Pokemon",
			function:		exploreRegion,
		},
		"catch": {
			name:				"catch",
			description: "Catch a Pokemon",
			function:		catchPokemon,
		},
		"inspect": {
			name:				"inspect",
			description: "Inspect a Pokemon from your Pokedex",
			function:		inspectPokemon,
		},
		"pokedex": {
			name:				"pokedex",
			description: "Display all Pokemon in your Pokedex",
			function:		displayPokedex,
		},
	}
}

func exitPokedex(cfg *config, args ...string) error {
	return errors.New("exit")
}

func listCommands(cfg *config, args ...string) error {
	for _, cmd := range commands {

		if cmd.name == "" || cmd.description == "" {
			return errors.New("Error getting commands")
		}
		
		fmt.Println(cmd.name + ": " + cmd.description)
	}
	return nil
}

func displayMap(cfg *config, args ...string) error {
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

func displayMapb(cfg *config, args ...string) error {
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

func exploreRegion(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location provided")
	}
	locationName := args[0]
	fmt.Println("Exploring " + locationName + "...")

	location, err := cfg.pokeapiClient.GetLocationPokemon(locationName)
	if err != nil {
		return fmt.Errorf("Error getting pokemon: %w", err)
	}
	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Println(pokemonEncounter.Pokemon.Name)
	}
	return nil
}

func catchPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No Pokemon provided")
	}
	pokemonName := args[0]
	fmt.Println("Catching " + pokemonName + "...")
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("Error getting pokemon: %w", err)
	}

	const threshold = 50
	// handle catching the pokemon
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum < threshold {
		fmt.Println("You caught " + pokemon.Name + "!")
		userPokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println("Oh no! " + pokemon.Name + " got away!")
	}

	return nil
}

func inspectPokemon(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No Pokémon provided. Usage: inspectPokemon <pokemon-name>")
	}

	pokemonName := strings.ToLower(args[0])
	pokemon, exists := userPokedex[pokemonName]
	if !exists {
		fmt.Println("Pokémon not found in Pokedex")
		return nil
	}

	fmt.Printf("Name: %s\n", strings.Title(pokemon.Name))
	fmt.Printf("ID: %d\n", pokemon.ID)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println()

	fmt.Println("Abilities:")
	for _, ability := range pokemon.Abilities {
		hidden := ""
		if ability.IsHidden {
			hidden = " (Hidden)"
		}
		fmt.Printf("- %s%s (Slot %d)\n", strings.Title(ability.Ability.Name), hidden, ability.Slot)
	}
	fmt.Println()

	fmt.Println("Types:")
	var types []string
	for _, t := range pokemon.Types {
		types = append(types, strings.Title(t.Type.Name))
	}
	fmt.Printf("- %s\n", strings.Join(types, ", "))
	fmt.Println()

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d (Effort: %d)\n", strings.Title(stat.Stat.Name), stat.BaseStat, stat.Effort)
	}
	fmt.Println()

	if len(pokemon.Forms) > 1 {
		fmt.Println("Forms:")
		for _, form := range pokemon.Forms {
			fmt.Printf("- %s\n", strings.Title(form.Name))
		}
		fmt.Println()
	}

	return nil
}

func displayPokedex(cfg *config, args ...string) error {
	if len(userPokedex) == 0 {
		fmt.Println("Pokedex is empty")
		return nil
	}

	for name := range userPokedex {
		fmt.Println(strings.Title(name))
	}

	return nil
}