package api

import (
	"cookbook/api/handlers"

	"github.com/gin-gonic/gin"
)

func InitAPI() *gin.Engine {
	api := gin.Default()
	api.Use(handlers.ErrorHandler())

	v1 := api.Group("/api/v1")
	v1.Any("/recipes", handlers.HandleRecipes)
	v1.GET("/recipes/:id", handlers.GetRecipe)
	v1.DELETE("/recipes/:id", handlers.DeleteRecipe)

	return api
}
