package user

import (
	"context"

	"github.com/fajarcandraaa/dating_app_be/internal/entity/account"
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
)

type UserRepositoryContract interface {
	SignIn(ctx context.Context, payload userPresentation.LoginRequest) (*userPresentation.User, error)
	SignUp(ctx context.Context, payload userEntity.User, payloadAccount account.Account) error

	GetRandom(ctx context.Context, gender userEntity.UserGender, userCode string, execptionUserCode []string) (*userEntity.User, error)
}
