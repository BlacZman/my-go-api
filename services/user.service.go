package services

import (
	"my-go-api/models"
)

type UserService struct {
	databaseService DatabaseService
}

func NewUserService(databaseService DatabaseService) UserService {
	return UserService{
		databaseService: databaseService,
	}
}

func (s UserService) GetUser(id uint) models.User {
	var user models.User
	s.databaseService.db.First(&user, id)
	return user
}
