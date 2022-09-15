package main

import (
	"github.com/RAMESSESII2/test-go-gin/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", handlers.ShowIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", handlers.GetArticle)

	router.Run()

}
