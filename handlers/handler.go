package handlers

import (
	"net/http"
	"strconv"

	"github.com/RAMESSESII2/test-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ShowIndexPage(ctx *gin.Context) {
	// Get all articles
	articles := models.GetAllArtciles()

	render(ctx, gin.H{
		"title":   "Home Page",
		"payload": articles,
	},
		"index.html",
	)
	// Not needed since render() covers it and more
	// call the HTML method of the Context to render a template
	// ctx.HTML(
	// 	// Set the HTTP status to 200 (OK)
	// 	http.StatusOK,
	// 	"index.html",
	// 	// Pass teh data that the page uses
	// 	gin.H{
	// 		"title":   "Home Page",
	// 		"payload": articles,
	// 	},
	// )
}

func GetArticle(ctx *gin.Context) {
	// check if article ID is valid
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		// check if the article exists
		if article, err := models.GetArticleById(articleID); err == nil {
			render(ctx, gin.H{
				"title":   article.Title,
				"payload": article,
			},
				"article.html",
			)
		} else {
			// if article is not found, abort with an error
			ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
		// if an invalid article ID is specified in the URL, log error, redirect to home page
		// log.Println("[ERROR]: ", http.StatusNotFound, " Couldn't find the article ID")
		// ctx.Redirect(http.StatusNotFound, "/")
	}
}
