package delivery

import (
	"golang-simple-api/internal/delivery/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(app *gin.Engine, handler *handler.SimpleHandler) {
	app.GET("/persons", handler.GetAll)
	// app.POST("/person", handler.Create)
	// app.GET("/person/:id", handler.GetByID)
}
