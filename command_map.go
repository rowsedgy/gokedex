package main

import (
	"fmt"

	"github.com/rowsedgy/gokedex/internal/pokeapi"
)

var id int = 1

var inc int = 20

func commandMap(cfg *Config) error {
	if cfg.Next != nil {
		cfg.Previous = cfg.Next
	}
	cfg.Next = nil
	for i := id; i < id+inc; i++ {
		name := pokeapi.GetLocationName(i)
		cfg.Next = append(cfg.Next, name)
	}
	id += inc
	for _, v := range cfg.Next {
		fmt.Println(v)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	if id <= 1 {
		fmt.Println("you're on the first page")
		return nil
	}
	id -= inc
	cfg.Next = nil

	for i := id; i < id+inc; i++ {
		name := pokeapi.GetLocationName(i)
		cfg.Next = append(cfg.Next, name)
	}

	if id-inc >= 1 {
		cfg.Previous = nil
		for i := id - inc; i < id; i++ {
			name := pokeapi.GetLocationName(i)
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
