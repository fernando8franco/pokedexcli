package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	val, isCached := c.cache.Get(url)
	if isCached {
		return unmarshalPokemonData(val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)

	return unmarshalPokemonData(data)
}

func unmarshalPokemonData(data []byte) (Pokemon, error) {
	var pokemonResp Pokemon
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
