package pokeapi

import (
	"encoding/json"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	var locationsResp Locations
	if err := json.NewDecoder(resp.Body).Decode(&locationsResp); err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
