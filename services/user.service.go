package services

import (
	"my-go-api/models"
)

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

func (u *UserService) getUser() models.User {
	return models.User{}
}
