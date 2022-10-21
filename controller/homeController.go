package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomepageHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message" : "Here is HomePage",
	})
}