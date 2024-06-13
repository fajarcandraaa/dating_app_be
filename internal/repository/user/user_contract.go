package user

import (
	"context"

	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
)

type UserRepositoryContract interface {
	SignIn(ctx context.Context, payload userPresentation.LoginRequest) (*userPresentation.User, error)
}