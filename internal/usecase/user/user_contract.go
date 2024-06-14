package user

import (
	"context"

	// userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
)

type UserAuthUsecaseContract interface {
	LogIn(ctx context.Context, payload *userPresentation.LoginRequest) (*string, error)
	SignUp(ctx context.Context, paylaod *userPresentation.RegistrationRequest) error
}

type UserInfoUsecaseContract interface {
	GetRandomProfile(ctx context.Context, userCode string) (*userPresentation.User, error)
}
