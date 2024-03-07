package services

import (
	"my-go-api/models"

	"gorm.io/gorm"
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

type UpdateUserBody struct {
	ID       uint   `json:"id" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s UserService) UpdateUser(updateUser UpdateUserBody) (*models.User, error) {
	user := models.User{
		Email:    updateUser.Email,
		Password: updateUser.Password,
		Model: gorm.Model{
			ID: updateUser.ID,
		},
	}

	if _, err := s.GetUser(updateUser.ID); err != nil {
		return nil, err
	}

	if result := s.databaseService.db.Model(&user).Where("id = ?", updateUser.ID).Updates(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s UserService) DeleteUser(id uint) (*uint, error) {
	if _, err := s.GetUser(id); err != nil {
		return nil, err
	}

	if result := s.databaseService.db.Delete(&models.User{}, id); result.Error != nil {
		return nil, result.Error
	}
	return &id, nil
}
