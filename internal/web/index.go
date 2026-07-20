package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {

	data, err := LoadIndexData()
	if err != nil {
		c.String(
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	c.HTML(
		http.StatusOK,
		"index.html",
		data,
	)
}
