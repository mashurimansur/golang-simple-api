package handler

import (
	"golang-simple-api/internal/model"
	"golang-simple-api/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PokemonHandler struct {
	PokemonUsecase usecase.PokemonUsecaseInterface
}

func NewPokemonHandler(pokemonUsecase usecase.PokemonUsecaseInterface) *PokemonHandler {
	return &PokemonHandler{
		PokemonUsecase: pokemonUsecase,
	}
}

func (p *PokemonHandler) GetPokemons(ctx *gin.Context) {
	pokemons, err := p.PokemonUsecase.GetPokemons()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Custom-Header", "test header")
	ctx.JSON(http.StatusOK, model.BaseResponse[[]model.PokemonsResponse]{
		Data: pokemons,
		Code: http.StatusOK,
	})
}
