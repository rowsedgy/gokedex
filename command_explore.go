package main

import (
	"fmt"

	"github.com/rowsedgy/gokedex/internal/pokeapi"
)

func commandExplore(cfg *Config, areaToExplore string) error {
	pokemonsInArea, err := pokeapi.GetPokemonNames(areaToExplore)
	if err != nil {
		fmt.Println("Error getting pokemons in area: ", err)
		return err
	}

	fmt.Printf("Exploring %s...\n", areaToExplore)
	for _, pkmn := range pokemonsInArea {
		fmt.Printf("- %s\n", pkmn)
	}
	return nil
}
