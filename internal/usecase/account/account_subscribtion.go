package account

import (
	"context"

	"github.com/fajarcandraaa/dating_app_be/internal/presentation/subscription"
	"github.com/fajarcandraaa/dating_app_be/internal/repository"
	"github.com/google/uuid"
)

type accountSubscribeUsecase struct {
	repo *repository.Repositories
}

func NewAccountSubscribeUsecase(repo *repository.Repositories) *accountSubscribeUsecase {
	return &accountSubscribeUsecase{repo}
}

var _ AccountSubscribtionUsecaseContract = &accountSubscribeUsecase{}

// Subscribe implements AccountSubscribtionUsecaseContract.
func (u *accountSubscribeUsecase) Subscribe(ctx context.Context, payload subscription.UpdatePackageRequest) error {
	var (
		userId = uuid.MustParse(payload.UserId)
	)
	err := u.repo.SubscribtionAccopunt.UpdatePackage(ctx, userId, payload.FeatureCode, payload.PackageId)
	if err != nil {
		return err
	}

	return nil
}
