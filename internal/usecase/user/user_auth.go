package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/fajarcandraaa/dating_app_be/config/auth"
	"github.com/fajarcandraaa/dating_app_be/internal/dto"
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type authUseCase struct {
	repo *repository.Repositories
	rds  *redis.Client
}

func NewAuthUsecase(repo *repository.Repositories, rds *redis.Client) *authUseCase {
	return &authUseCase{
		repo: repo,
		rds:  rds,
	}
}

var _ UserAuthUsecaseContract = &authUseCase{}

// Login implements UserUsecaseContract.
func (u *authUseCase) LogIn(ctx context.Context, payload *userPresentation.LoginRequest) (*string, error) {
	var (
		expiration = 1 * 24 * time.Hour
	)
	singIn, err := u.repo.User.SignIn(ctx, *payload)
	if err != nil {
		return nil, err
	}

	getToken, err := auth.CreateToken(singIn.UserCode, singIn.Id)
	if err != nil {
		return nil, errors.Wrap(err, "build token")
	}

	go func() {
		rdsKey := fmt.Sprintf("login_%s", singIn.UserCode)
		dataStore, err := json.Marshal(singIn)
		if err != nil {
			log.Printf("ERROR Redis Set Payload : %s", err.Error())
		}
		err = u.rds.Set(context.Background(), rdsKey, dataStore, expiration).Err()
		if err != nil {
			log.Printf("ERROR Redis Set : %s", err.Error())
		}
	}()

	return &getToken, nil
}

// SignUp implements UserAuthUsecaseContract.
func (u *authUseCase) SignUp(ctx context.Context, payload *userPresentation.RegistrationRequest) error {
	userPayload, accountPayload, err := dto.TransforRegistrationToDatabase(*payload)
	if err != nil {
		return err
	}
	err = userEntity.UserRequestValidate(userPayload)
	if err != nil {
		return err
	}

	err = u.repo.User.SignUp(ctx, *userPayload, *accountPayload)
	if err != nil {
		return err
	}

	return nil
}
