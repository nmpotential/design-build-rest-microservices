package services

import "github.com/nmpotential/design-build-rest-microservices/bookstore_users-api/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
