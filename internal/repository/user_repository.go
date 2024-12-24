package repository

import (
	"go-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	GetAllUsers() ([]model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user *model.User) error {

	return u.db.Create(user).Error
}

func (u *userRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User

	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) UpdateUser(user *model.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) DeleteUser(id uint) error {
	return u.db.Delete(&model.User{}, id).Error
}
