package user

import (
	"context"

	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
)

type UserAuthUsecaseContract interface {
	LogIn(ctx context.Context, payload *userPresentation.LoginRequest) (*string, error)
	//TODO :
	// 1. Create SingUp function
	// 2. Create LogOut function
}