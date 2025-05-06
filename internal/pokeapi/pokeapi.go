package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func GetNextLocation(cfg *Config) error {
	res, err := http.Get(cfg.NextUrl)
	if err != nil {
		return fmt.Errorf("Error on HTTP GET: %w", err)
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading JSON response: %w", err)
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

func GetPrevLocation(cfg *Config) error {
	if cfg.PrevUrl == "" {
		fmt.Println("you're on the first page")
	}
	res, err := http.Get(cfg.PrevUrl)
	if err != nil {
		return fmt.Errorf("Error on HTTP GET: %w", err)
	}
	defer res.Body.Close()

	jsonData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("Error reading JSON response: %w", err)
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
