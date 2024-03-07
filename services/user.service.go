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

func (s UserService) GetUser(id uint) (*models.User, error) {
	var user models.User

	if result := s.databaseService.db.First(&user, id); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

type CreateUserBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s UserService) CreateUser(newUser CreateUserBody) (*models.User, error) {
	user := models.User{
		Email:    newUser.Email,
		Password: newUser.Password,
	}

	if result := s.databaseService.db.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
