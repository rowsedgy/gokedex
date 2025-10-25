package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rowsedgy/gokedex/internal/pokecache"
)

var baseURL string = "https://pokeapi.co/api/v2/location-area/"

var locationCache = pokecache.NewCache(10 * time.Second)

var pokemonCache = pokecache.NewCache(10 * time.Second)

func httpReq(url string) (*http.Response, error) {
	// url debugging
	// fmt.Println("URL IN REQUEST: ", url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("WrongStatusCode: %d", res.StatusCode)
	}

	return res, nil
}

func GetLocationName(id int) (string, error) {
	cacheKey := fmt.Sprintf("%d", id)

	if val, ok := locationCache.Get(cacheKey); ok {
		return string(val), nil
	}

	idURL := fmt.Sprintf("%s%d/", baseURL, id)

	res, err := httpReq(idURL)
	if res == nil {
		return "", err
	}

	defer res.Body.Close()

	var locationData LocationArea
	if err := json.NewDecoder(res.Body).Decode(&locationData); err != nil {
		fmt.Println("Error decoding json", err)
		return "", nil
	}

	locationCache.Add(cacheKey, []byte(locationData.Name))
	return locationData.Name, nil

}

func GetPokemonNames(areaName string) ([]string, error) {
	if val, ok := pokemonCache.Get(areaName); ok {
		// CACHE HIT MESSAGE
		// fmt.Println("CACHE HIT!!")
		var pkmn []string
		if err := json.Unmarshal(val, &pkmn); err != nil {
			fmt.Println("Error unmarshaling cache hit", err)
			return nil, err
		}
		return pkmn, nil
	}

	areaURL := fmt.Sprintf("%s%s/", baseURL, areaName)

	res, err := httpReq(areaURL)
	if res == nil {
		return nil, err
	}
	defer res.Body.Close()

	var areaData LocationArea

	if err := json.NewDecoder(res.Body).Decode(&areaData); err != nil {
		fmt.Println("Error decoding json", err)
		return nil, err
	}

	pkmn := []string{}

	if _, ok := pokemonCache.Get(areaName); !ok {
		// ADDING TO CACHE MESSAGE
		// fmt.Printf("ADDING %s TO CACHE", areaName)
		for _, encounter := range areaData.PokemonEncounters {
			pkmn = append(pkmn, encounter.Pokemon.Name)
		}
		encodedPkmn, err := json.Marshal(pkmn)
		if err != nil {
			fmt.Println("Error marshaling json", err)
			return nil, err
		}
		pokemonCache.Add(areaName, encodedPkmn)
	}

	return pkmn, nil

}
