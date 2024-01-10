package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"todo-api/docs"
	"todo-api/handlers"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("api/v1")
	{
		v1.GET("/todos", handlers.GetTodos)
	}
	docs.SwaggerInfo.BasePath = "api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
