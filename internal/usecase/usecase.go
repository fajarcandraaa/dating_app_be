package usecase

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
	"github.com/go-redis/redis/v8"
)

type UsecaseContract struct {
	UserAuthContract user.UserAuthUsecaseContract
	UserInfoContract user.UserInfoUsecaseContract
}

func NewUsecaseService(repo *repository.Repositories, rds *redis.Client) *UsecaseContract {
	userAuthUsecase := user.NewAuthUsecase(repo, rds)
	userInfoUsecase := user.NewUserInfoUsecase(repo, rds)

	return &UsecaseContract{
		UserAuthContract: userAuthUsecase,
		UserInfoContract: userInfoUsecase,
	}
}
