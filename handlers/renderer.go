package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func render(ctx *gin.Context, data gin.H, templateName string) {
	switch ctx.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSOn
		ctx.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with JSOn
		ctx.XML(http.StatusOK, data["payload"])
	default:
		ctx.HTML(http.StatusOK, templateName, data)
	}
}
