package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/", indexHandler)

	router.POST("/upload", handlers.UploadHash)
}
