package user

import (
	"context"

	"github.com/fajarcandraaa/dating_app_be/config/auth"
	"github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/pkg/errors"
)

type authUseCase struct {
	repo *repository.Repositories
}

func NewAuthUsecase(repo *repository.Repositories) *authUseCase {
	return &authUseCase{
		repo: repo,
	}
}

var _ UserAuthUsecaseContract = &authUseCase{}

// Login implements UserUsecaseContract.
func (u *authUseCase) LogIn(ctx context.Context, payload *user.LoginRequest) (*string, error) {
	singIn, err := u.repo.User.SignIn(ctx, *payload)
	if err != nil {
		return nil, err
	}

	getToken, err := auth.CreateToken(singIn.UserCode, singIn.Id)
	if err != nil {
		return nil, errors.Wrap(err, "build token")
	}

	//TODO : Have to create a function to strore user data into redis

	return &getToken, nil
}
