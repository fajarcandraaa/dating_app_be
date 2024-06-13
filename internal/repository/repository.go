package repository

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User user.UserRepositoryContract
}
func NewRepositories(db *gorm.DB) *Repositories {
	userRepo := user.NewUserRepository(db)
	
	return &Repositories{
		User: userRepo,
	}
}
