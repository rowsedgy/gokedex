package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/rowsedgy/gokedex/internal/pokecache"
)

var locationCache = pokecache.NewCache(10 * time.Second)

type Location struct {
	Name string `json:"name"`
}

func GetLocationName(id int) string {
	key := fmt.Sprintf("%d", id)

	if val, ok := locationCache.Get(key); ok {
		// fmt.Println("cache hit")
		return string(val)
	}

	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", id)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("HTTP Error:", res.StatusCode)
		return ""
	}

	var data Location
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println("JSON decode error:", err)
		return ""
	}

	locationCache.Add(key, []byte(data.Name))
	return data.Name
}
