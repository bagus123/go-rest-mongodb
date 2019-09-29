package user

import (
	"net/http"

	"github.com/bagus123/go-rest-mongodb/models"
	"github.com/bagus123/go-rest-mongodb/services"
	"github.com/bagus123/go-rest-mongodb/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var userService = new(services.UserService)

// User ...
type User = models.User

// Response : alias of r.Response
type Response = response.Response

// ResponseError : alias of r.Error
type ResponseError = response.Error

// GetAllUsers ...
func GetAllUsers(c *gin.Context) {
	res, err := userService.GetAllUsers()
	if err == nil {
		response := &Response{
			Data: map[string]interface{}{
				"users": res,
			},
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(400, &ResponseError{Status: 400, Message: err.Error()})
	}

}

// GetUserByID ...
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	res, err := userService.GetUserByID(id)
	if err == nil {
		response := &Response{
			Data: map[string]interface{}{
				"user": res,
			},
		}
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(400, &ResponseError{Status: 400, Message: err.Error()})
	}
}

// UpdateUserByID ...
func UpdateUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		c.JSON(400, &ResponseError{Status: 400, Message: err.Error()})
	}

	res, err := userService.UpdateUserByID(id, user)
	if err == nil {
		response := &Response{
			Data: map[string]interface{}{
				"user": res,
			},
		}
		c.JSON(http.StatusOK, response)
	} else {
		// status 400 also available in constant list : http.StatusBadRequest
		c.JSON(400, &ResponseError{Status: 400, Message: err.Error()})
	}
}
