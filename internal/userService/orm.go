package userService

import (
	"project/internal/taskService"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Task     []taskService.Task `gorm:"foreignKey:UserID" json:"task"`
}
