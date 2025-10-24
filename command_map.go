package main

import (
	"fmt"

	"github.com/rowsedgy/gokedex/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	if cfg.NoMoreLocations {
		fmt.Println("No more locations!")
		return nil
	}

	if cfg.Next != nil {
		cfg.Previous = cfg.Next
	}

	cfg.Next = nil

	for i := cfg.ID; i < cfg.ID+cfg.Increment; i++ {
		name, err := pokeapi.GetLocationName(i)
		if err != nil {
			fmt.Println("No more locations!")
			cfg.NoMoreLocations = true
			break
		}
		cfg.Next = append(cfg.Next, name)
	}

	for _, v := range cfg.Next {
		fmt.Println(v)
	}

	if cfg.NoMoreLocations || len(cfg.Next) < cfg.Increment {
		fmt.Println("No more locations!")
	}

	cfg.ID += cfg.Increment

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.ID <= 1 {
		fmt.Println("you're on the first page")
		return nil
	}

	cfg.ID -= cfg.Increment
	cfg.Next = nil

	for i := cfg.ID; i < cfg.ID+cfg.Increment; i++ {
		name, err := pokeapi.GetLocationName(i)
		if err != nil {
			cfg.NoMoreLocations = true
			break
		}
		cfg.Next = append(cfg.Next, name)
	}

	if cfg.ID-cfg.Increment >= 1 {
		cfg.Previous = nil
		for i := cfg.ID - cfg.Increment; i < cfg.ID; i++ {
			name, err := pokeapi.GetLocationName(i)
			if err != nil {
				cfg.NoMoreLocations = true
				break
			}
			cfg.Previous = append(cfg.Previous, name)
		}
	} else {
		cfg.Previous = nil
		fmt.Println("you're on the first page")
	}

	for _, v := range cfg.Previous {
		fmt.Println(v)
	}

	return nil
}
