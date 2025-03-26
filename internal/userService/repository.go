package userService

import (
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		if strings.Contains(err.Error(), `relation "users" does not exist`) {
			return []User{}, nil
		}
		return nil, err
	}
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.Password != "" {
		existingUser.Password = user.Password
	}

	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
