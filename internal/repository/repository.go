package repository

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository/account"
	"github.com/fajarcandraaa/dating_app_be/internal/repository/user"
	"gorm.io/gorm"
)

type Repositories struct {
	User                 user.UserRepositoryContract
	SubscribtionAccopunt account.AccountRepositoryContract
}

func NewRepositories(db *gorm.DB) *Repositories {
	userRepo := user.NewUserRepository(db)
	subscribeAccountRepo := account.NewAccountRepository(db)

	return &Repositories{
		User:                 userRepo,
		SubscribtionAccopunt: subscribeAccountRepo,
	}
}
