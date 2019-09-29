package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllPosts ...
func GetAllPosts(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
