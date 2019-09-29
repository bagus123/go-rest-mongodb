package services

import (
	"github.com/bagus123/go-rest-mongodb/daos"
	"github.com/bagus123/go-rest-mongodb/models"

	"github.com/pkg/errors"
)

// User : alias of models.User
type User = models.User

// UserService ...
type UserService struct{}

// GetAllUsers ...
func (s *UserService) GetAllUsers() ([]interface{}, error) {
	results, err := daos.UserDAO().GetAll()

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return results, nil
}

// GetUserByID ...
func (s *UserService) GetUserByID(id string) (interface{}, error) {
	result, err := daos.UserDAO().GetOneByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return result, nil
}

// UpdateUserByID ...
func (s *UserService) UpdateUserByID(id string, user User) (interface{}, error) {

	currentUser, err := daos.UserDAO().GetOneByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	// casting from interface{} to User
	updateData, ok := currentUser.(User)

	if ok {
		updateData.UserName = user.UserName
		updateData.Password = user.Password
		updateData.Role = user.Role
	}

	result, err := daos.UserDAO().UpdateOneByID(id, updateData)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	result, err = daos.UserDAO().GetOneByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return result, nil
}
