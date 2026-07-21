package web

import "github.com/gin-gonic/gin"

func apiQueue(c *gin.Context) {
	c.JSON(200, listQueueJobs())
}
