package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gyunn35/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(t time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(t),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocationAreas(url *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if url != nil {
		fullURL = *url
	}

	if v, ok := c.cache.Get(fullURL); ok {
		fmt.Println("cache hit!")
		locationAreaResponse := LocationAreaResponse{}
		err := json.Unmarshal(v, &locationAreaResponse)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		return locationAreaResponse, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	if response.StatusCode >= 400 {
		return LocationAreaResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer response.Body.Close()

	locationAreaResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResponse, nil
}

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	endpoint := "/location-area/" + name
	fullURL := baseURL + endpoint

	if v, ok := c.cache.Get(fullURL); ok {
		fmt.Println("cache hit!")
		locationAreaResponse := LocationArea{}
		err := json.Unmarshal(v, &locationAreaResponse)
		if err != nil {
			return LocationArea{}, err
		}
		return locationAreaResponse, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	if response.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return LocationArea{}, err
	}
	defer response.Body.Close()

	locationAreaResponse := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResponse)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreaResponse, nil
}

func (c *Client) GetPokemon(name string) (PokemonResponse, error) {
	endpoint := "/pokemon/" + name
	fullURL := baseURL + endpoint

	if v, ok := c.cache.Get(fullURL); ok {
		fmt.Println("cache hit!")
		pokemonResponse := PokemonResponse{}
		err := json.Unmarshal(v, &pokemonResponse)
		if err != nil {
			return PokemonResponse{}, err
		}
		return pokemonResponse, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	if response.StatusCode >= 400 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer response.Body.Close()

	pokemonResponse := PokemonResponse{}
	err = json.Unmarshal(data, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonResponse, nil
}
