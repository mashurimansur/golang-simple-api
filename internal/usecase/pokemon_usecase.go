package usecase

import (
	"golang-simple-api/internal/model"
	"golang-simple-api/internal/provider"
)

type PokemonUsecase struct {
	PokemonClient provider.PokemonClient
}

func NewPokemonUsecase(pokemonClient provider.PokemonClient) *PokemonUsecase {
	return &PokemonUsecase{
		PokemonClient: pokemonClient,
	}
}

type PokemonUsecaseInterface interface {
	GetPokemons() ([]model.PokemonsResponse, error)
}

func (p *PokemonUsecase) GetPokemons() ([]model.PokemonsResponse, error) {
	pokemonResponse, err := p.PokemonClient.GetPokemons()
	if err != nil {
		return nil, err
	}

	return pokemonResponse.Results, nil
}
