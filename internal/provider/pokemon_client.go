package provider

import (
	"fmt"
	"golang-simple-api/internal/model"
	"golang-simple-api/pkg/httpclient"
)

type PokemonClientInterface interface {
	GetPokemons() (model.PokemonResponse, error)
}

type PokemonClient struct {
	BaseUrl    string
	httpClient *httpclient.Client
}

func NewPokemonClient(httpClient *httpclient.Client, baseUrl string) *PokemonClient {
	return &PokemonClient{
		httpClient: httpClient,
		BaseUrl:    baseUrl,
	}
}

func (p *PokemonClient) GetPokemons() (model.PokemonResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}

	fmt.Println(p.BaseUrl)
	url := p.BaseUrl + "/api/v2/pokemon"

	res, err := p.httpClient.DoRequest("GET", url, headers, nil)
	if err != nil {
		return model.PokemonResponse{}, err
	}

	var pokemonResponse model.PokemonResponse
	err = httpclient.ReadResponse(res, &pokemonResponse)
	if err != nil {
		return model.PokemonResponse{}, err
	}

	return pokemonResponse, nil
}
