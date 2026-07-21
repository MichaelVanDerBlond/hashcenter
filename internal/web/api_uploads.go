package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func apiUploads(c *gin.Context) {

	uploads, err := LoadUploads()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, uploads)
}
