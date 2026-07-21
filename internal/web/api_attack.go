package web

import (
	"net/http"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
	"github.com/gin-gonic/gin"
)

type AttackRequest struct {
	HashFile     string   `json:"hash_file"`
	Dictionaries []string `json:"dictionaries"`

	HashMode   int `json:"hash_mode"`
	AttackMode int `json:"attack_mode"`

	Rule        string `json:"rule"`
	Mask        string `json:"mask"`
	SessionName string `json:"session_name"`
	Device      string `json:"device"`
	Workload    int    `json:"workload"`

	Extra []string `json:"extra"`
}

func apiAttack(c *gin.Context) {

	var req AttackRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(req.Dictionaries) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "no dictionaries selected",
		})
		return
	}

	for _, dict := range req.Dictionaries {

		job := attack.Job{
			HashFile:    req.HashFile,
			Dictionary:  dict,
			HashMode:    req.HashMode,
			AttackMode:  req.AttackMode,
			Rule:        req.Rule,
			Mask:        req.Mask,
			SessionName: req.SessionName,
			Device:      req.Device,
			Workload:    req.Workload,
			Extra:       append([]string(nil), req.Extra...),
		}

		enqueue(job)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status": "queued",
		"jobs":   len(req.Dictionaries),
	})
}
