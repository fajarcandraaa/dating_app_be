package account

import (
	"context"

	"github.com/fajarcandraaa/dating_app_be/internal/presentation/subscription"
)

type AccountSubscribtionUsecaseContract interface {
	Subscribe(ctx context.Context, payload subscription.UpdatePackageRequest) error
}