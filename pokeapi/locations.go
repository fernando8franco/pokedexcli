package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, isCached := c.cache.Get(url)
	if isCached {
		return unmarshalLocationData(val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}
	c.cache.Add(url, data)

	return unmarshalLocationData(data)
}

func (c *Client) ListPokemonsInLocation(location string) (PokemonsByLocation, error) {
	url := baseURL + "/location-area/" + location

	val, isCached := c.cache.Get(url)
	if isCached {
		return unmarshalPokemonLocationData(val)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonsByLocation{}, err
	}

	resp, err := c.httClient.Do(req)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return PokemonsByLocation{}, nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonsByLocation{}, err
	}
	c.cache.Add(url, data)

	return unmarshalPokemonLocationData(data)
}

func unmarshalLocationData(data []byte) (Locations, error) {
	var locationsResp Locations
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}

func unmarshalPokemonLocationData(data []byte) (PokemonsByLocation, error) {
	var locationsResp PokemonsByLocation
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return PokemonsByLocation{}, err
	}

	return locationsResp, nil
}
