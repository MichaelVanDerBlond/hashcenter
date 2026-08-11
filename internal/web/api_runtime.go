package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/services"
	"github.com/gin-gonic/gin"
)

func apiRuntime(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		services.ListRuntime(),
	)
}
