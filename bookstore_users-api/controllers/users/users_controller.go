package users

import (
	"github.com/gin-gonic/gin"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/domain/users"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/services"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		// Returns bad request to the caller.
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// Handles user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
