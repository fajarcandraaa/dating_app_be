package account

import (
	"context"

	"github.com/google/uuid"
)

type AccountRepositoryContract interface {
	UpdatePackage(ctx context.Context, userId uuid.UUID, featureCode string, packageId int) error
}
