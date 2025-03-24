package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"date":    time.Now(),
	})
}
