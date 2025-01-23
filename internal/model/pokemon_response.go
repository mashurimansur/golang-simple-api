package model

type PokemonResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []PokemonsResponse
}

type PokemonsResponse struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
