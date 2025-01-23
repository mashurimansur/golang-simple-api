package app

import (
	"golang-simple-api/internal/delivery"
	"golang-simple-api/internal/delivery/handler"
	"golang-simple-api/internal/provider"
	"golang-simple-api/internal/repository"
	"golang-simple-api/internal/usecase"
	"golang-simple-api/pkg/config"
	"golang-simple-api/pkg/db"
	"golang-simple-api/pkg/httpclient"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() error {
	cfg := config.LoadConfig()
	dbConn := db.NewMySQLConnection(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	err := dbConn.Connect()
	if err != nil {
		log.Fatal("Error connecting to database ", err.Error())
		return err
	}
	defer dbConn.Close()

	simpleRepository := repository.NewSimpleRepository(dbConn.DB)
	simpleUsecase := usecase.NewSimpleUsecase(simpleRepository)
	simpleHandler := handler.NewSimpleHandler(simpleUsecase)

	httpClient := httpclient.NewClient(time.Second * 10)
	pokemonProvider := provider.NewPokemonClient(httpClient, cfg.PokemonAPI)
	pokemonUsecase := usecase.NewPokemonUsecase(*pokemonProvider)
	pokemonHandler := handler.NewPokemonHandler(pokemonUsecase)

	app := gin.Default()
	delivery.RegisterRoutesPerson(app, simpleHandler)
	delivery.RegisterRoutesPokemon(app, pokemonHandler)

	return app.Run(":8080")
}
