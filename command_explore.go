package main

import (
	"time"

	"github.com/rowsedgy/gokedex/internal/pokecache"
)

var exploreCache = pokecache.NewCache(10 * time.Second)

type LocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

// func commandExplore(cfg *Config) error {

// }
