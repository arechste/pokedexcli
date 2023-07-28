package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon" + "/" + pokemonName
	fullURL := baseUrl + endpoint

	// check the cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		fmt.Println("cache hit!")
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return Pokemon{}, nil
	}
	fmt.Println("cache miss!")

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, err
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}
	// cache miss polulate cache before return resp
	//fmt.Println(string(dat))
	c.cache.Add(fullURL, dat)
	return pokemon, nil
}
