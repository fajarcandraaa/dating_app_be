package usecase

import (
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/account"
	"github.com/fajarcandraaa/dating_app_be/internal/usecase/user"
	"github.com/go-redis/redis/v8"
)

type UsecaseContract struct {
	UserAuthContract user.UserAuthUsecaseContract
	UserInfoContract user.UserInfoUsecaseContract
	AccountSubscribtionContract account.AccountSubscribtionUsecaseContract
}

func NewUsecaseService(repo *repository.Repositories, rds *redis.Client) *UsecaseContract {
	userAuthUsecase := user.NewAuthUsecase(repo, rds)
	userInfoUsecase := user.NewUserInfoUsecase(repo, rds)
	accountSubscribeUsecase := account.NewAccountSubscribeUsecase(repo)

	return &UsecaseContract{
		UserAuthContract:            userAuthUsecase,
		UserInfoContract:            userInfoUsecase,
		AccountSubscribtionContract: accountSubscribeUsecase,
	}
}
