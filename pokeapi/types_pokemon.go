package pokeapi

type Pokemon struct {
	Id             int    `json:"count"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Types          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"types"`
}
