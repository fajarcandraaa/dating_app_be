package account

import (
	"context"
	"time"

	"github.com/fajarcandraaa/dating_app_be/internal/entity/subscription"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

var _ AccountRepositoryContract = &AccountRepository{}

// UpdatePackage implements AccountRepositoryContract.
func (r *AccountRepository) UpdatePackage(ctx context.Context, userId uuid.UUID, featureCode string, packageId int) error {
	tx := r.db.Begin()
	payloadToUpdateUser := user.User{
		PackageId:   &packageId,
		FeatureCode: &featureCode,
	}
	err := tx.Model(user.User{}).Where("id = ?", userId).Updates(payloadToUpdateUser).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	currentTime := time.Now()
	expDay := currentTime.AddDate(0, 0, 30)
	subHistory := subscription.SubscribeHistory{
		UserId:    &userId,
		PackageId: &packageId,
		ExpiredAt: &expDay,
	}
	err = tx.Create(&subHistory).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
