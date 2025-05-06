package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"sort"
	"strings"
)

func cleanInput(text string) []string {
	var splitString []string
	lowerString := strings.ToLower(text)
	lowerString = strings.TrimSpace(lowerString)
	splitString = strings.Fields(lowerString)
	return splitString
}

func commandExit(config *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s: %s\n", name, commands[name].description)
	}
	return nil
}

func commandMap(config *pokeapi.Config) error {
	pokeapi.GetNextLocation(config)
	return nil
}

func commandMapb(config *pokeapi.Config) error {
	pokeapi.GetPrevLocation(config)
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func main() {
	config := &pokeapi.Config{
		NextUrl: "https://pokeapi.co/api/v2/location-area",
		PrevUrl: "",
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		commandName := cleanInput(text)
		if len(commandName) == 0 {
			continue
		}

		command, exists := commands[commandName[0]]
		if exists {
			err := command.callback(config)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
