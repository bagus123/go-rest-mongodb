package api

import (
	postController "github.com/bagus123/go-rest-mongodb/controllers/v1/api/post"
	userController "github.com/bagus123/go-rest-mongodb/controllers/v1/api/user"

	"github.com/gin-gonic/gin"
)

// Init ...
func Init(g *gin.RouterGroup) {

	// USER
	g.GET("/user", userController.GetAllUsers)
	g.GET("/user/:id", userController.GetUserByID)
	g.PUT("/user/:id", userController.UpdateUserByID)

	// POST
	g.GET("/post", postController.GetAllPosts)
}
