package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/handlers"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {

	router.GET("/", indexHandler)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/api/uploads", apiUploads)
	router.GET("/api/dictionaries", apiDictionaries)
	router.GET("/api/queue", apiQueue)
	router.GET("/api/runtime", apiRuntime)

	router.POST("/api/attack", apiAttack)

	router.POST("/upload", handlers.UploadHash)
}
