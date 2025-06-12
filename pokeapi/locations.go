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
		return unmarshalData(val)
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

	return unmarshalData(data)
}

func unmarshalData(data []byte) (Locations, error) {
	var locationsResp Locations
	if err := json.Unmarshal(data, &locationsResp); err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
