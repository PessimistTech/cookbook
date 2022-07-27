package api

import "github.com/gin-gonic/gin"

func InitAPI() *gin.Engine {
	api := gin.Default()

	api.Group("/api/v1/recipes")

	return api
}
