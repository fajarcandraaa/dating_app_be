package usecase

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
)

type UsecaseContract struct {
	UserAuthContract user.UserAuthUsecaseContract
}

func NewUsecaseService(repo *repository.Repositories) *UsecaseContract {
	userAuthUsecase := user.NewAuthUsecase(repo)

	return &UsecaseContract{
		UserAuthContract: userAuthUsecase,
	}
}
