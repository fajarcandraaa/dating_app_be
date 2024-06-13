package usecase

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
	"github.com/go-redis/redis/v8"
)

type UsecaseContract struct {
	UserAuthContract user.UserAuthUsecaseContract
}

func NewUsecaseService(repo *repository.Repositories, rds *redis.Client) *UsecaseContract {
	userAuthUsecase := user.NewAuthUsecase(repo, rds)

	return &UsecaseContract{
		UserAuthContract: userAuthUsecase,
	}
}
