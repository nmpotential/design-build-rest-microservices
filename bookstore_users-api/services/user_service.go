package services

import (
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/domain/users"
	"github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, nil

	return &user, nil

	return &user, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}

}
