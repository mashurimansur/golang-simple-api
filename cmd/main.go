package main

import (
	"golang-simple-api/internal/app"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := app.NewApp().Start(); err != nil {
		log.Fatal("failed start: ", err.Error())
	}
}
