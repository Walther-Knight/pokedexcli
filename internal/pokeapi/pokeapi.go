package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/internal/pokecache"
)

type Config struct {
	NextUrl string
	PrevUrl string
}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Pokemon struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetNextLocation(cfg *Config, cache *pokecache.Cache, params string) error {
	jsonData, found := cache.Get(cfg.NextUrl)
	if !found {
		res, err := http.Get(cfg.NextUrl)
		if err != nil {
			return fmt.Errorf("Error on HTTP GET: %w", err)
		}
		defer res.Body.Close()
		var err2 error
		jsonData, err2 = io.ReadAll(res.Body)
		if err2 != nil {
			return fmt.Errorf("Error reading JSON response: %w", err2)
		}
		cache.Add(cfg.NextUrl, jsonData)
	}

	var location Location
	if err := json.Unmarshal(jsonData, &location); err != nil {
		return fmt.Errorf("Error Unmarshalling JSON: %w", err)
	}
	cfg.NextUrl = location.Next
	cfg.PrevUrl = location.Previous
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func GetPrevLocation(cfg *Config, cache *pokecache.Cache, params string) error {
	if cfg.PrevUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	jsonData, found := cache.Get(cfg.PrevUrl)
	if !found {
		res, err := http.Get(cfg.PrevUrl)
		if err != nil {
			return fmt.Errorf("Error on HTTP GET: %w", err)
		}
		defer res.Body.Close()

		var err2 error
		jsonData, err2 = io.ReadAll(res.Body)
		if err2 != nil {
			return fmt.Errorf("Error reading JSON response: %w", err2)
		}
		cache.Add(cfg.PrevUrl, jsonData)
	}

	var location Location
	if err := json.Unmarshal(jsonData, &location); err != nil {
		return fmt.Errorf("Error Unmarshalling JSON: %w", err)
	}
	cfg.NextUrl = location.Next
	cfg.PrevUrl = location.Previous
	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func GetPokemonInLocation(cfg *Config, cache *pokecache.Cache, params string) error {
	fullURL := cfg.NextUrl + "/" + params
	jsonData, found := cache.Get(fullURL)
	if !found {
		res, err := http.Get(fullURL)
		if err != nil {
			return fmt.Errorf("Error on HTTP GET: %w", err)
		}
		defer res.Body.Close()
		var err2 error
		jsonData, err2 = io.ReadAll(res.Body)
		if err2 != nil {
			return fmt.Errorf("Error reading JSON response: %w", err2)
		}
		cache.Add(cfg.NextUrl, jsonData)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(jsonData, &pokemon); err != nil {
		return fmt.Errorf("Error Unmarshalling JSON: %w", err)
	}
	for _, pok := range pokemon.PokemonEncounters {
		fmt.Println(pok.Pokemon.Name)
	}
	return nil
}
