package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/hashcatruntime"
	"github.com/gin-gonic/gin"
)

func apiRuntime(c *gin.Context) {

	c.JSON(
		http.StatusOK,
		hashcatruntime.List(),
	)
}
