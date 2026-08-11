package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
	"github.com/gin-gonic/gin"
)

func apiQueue(c *gin.Context) {
	c.JSON(http.StatusOK, task.DefaultManager().All())
}
