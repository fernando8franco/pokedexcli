package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

var URLNext = "https://pokeapi.co/api/v2/location-area/"
var URLPrevious *string

type config struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

type location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getMaps(back bool) (config, error) {
	var url string
	if !back {
		url = URLNext
	} else {
		if URLPrevious == nil {
			return config{}, errors.New("bad request")
		}

		url = *URLPrevious
	}

	res, err := http.Get(url)
	if err != nil {
		return config{}, err
	}
	defer res.Body.Close()

	var conf config
	if err := json.NewDecoder(res.Body).Decode(&conf); err != nil {
		return config{}, err
	}
	URLNext = *conf.Next
	URLPrevious = conf.Previous

	return conf, nil
}
