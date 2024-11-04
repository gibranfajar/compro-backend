package main

import (
	"compro-backend/config"
	"compro-backend/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database and migrate models
	config.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Success Connect API!",
		})
	})

	// route categories
	router.GET("/categories", controller.GetCategories)
	router.POST("/categories", controller.CreateCategory)
	router.GET("/categories/:id", controller.GetCategory)
	router.PUT("/categories/:id", controller.UpdateCategory)
	router.DELETE("/categories/:id", controller.DeleteCategory)

	// route articles
	router.GET("/articles", controller.GetArticles)
	router.POST("/articles", controller.CreateArticle)
	router.GET("/articles/:id", controller.GetArticle)
	router.PUT("/articles/:id", controller.UpdateArticle)
	router.DELETE("/articles/:id", controller.DeleteArticle)

	// Run the server on port 8080
	router.Run(":8080")

}
