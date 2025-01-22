package app

import (
	"golang-simple-api/internal/delivery"
	"golang-simple-api/internal/delivery/handler"
	"golang-simple-api/internal/repository"
	"golang-simple-api/internal/usecase"
	"golang-simple-api/pkg/config"
	"golang-simple-api/pkg/db"
	"log"

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

	app := gin.Default()
	delivery.RegisterRoutes(app, simpleHandler)

	return app.Run(":8080")
}
