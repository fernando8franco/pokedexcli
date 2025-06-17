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
		Type struct {
			Name string `json:"name"`
		}
	} `json:"types"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}
