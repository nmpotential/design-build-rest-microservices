package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/domain/users"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/services"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		fmt.Println(err)
		//TODO: return bad request to the caller.
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
