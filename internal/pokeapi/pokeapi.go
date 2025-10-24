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

// var pokemonCache = pokecache.NewCache(10 * time.Second)

type Location struct {
	Name string `json:"name"`
}

func httpReq(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Request error:", err)
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		// fmt.Printf("Non OK status code %d", res.StatusCode)
		res.Body.Close()
		return nil, fmt.Errorf("WrongStatusCode: %d", res.StatusCode)
	}

	return res, nil
}

func GetLocationName(id int) (string, error) {
	key := fmt.Sprintf("%d", id)

	if val, ok := locationCache.Get(key); ok {
		return string(val), nil
	}

	idURL := fmt.Sprintf("%s%d/", baseURL, id)

	res, err := httpReq(idURL)
	if res == nil {
		return "", err
	}

	defer res.Body.Close()

	var locationData Location
	if err := json.NewDecoder(res.Body).Decode(&locationData); err != nil {
		fmt.Println("Error decoding json", err)
		return "", nil
	}

	locationCache.Add(key, []byte(locationData.Name))
	return locationData.Name, nil

}
