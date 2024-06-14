package subscription

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubscribtionPackage int
type FeaturePackage string

const (
	PackageNonPremi      SubscribtionPackage = 1
	PackagePremium       SubscribtionPackage = 2
	FeatureNoQuota       FeaturePackage      = "NQA"
	FeatureVerifiedLabel FeaturePackage      = "VL"
)

type (
	Package struct {
		*gorm.DeletedAt

		Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;type:BIGINT"`
		Name      *string    `json:"name" gorm:"size:50;"`
		CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	PremiumFeature struct {
		*gorm.DeletedAt

		Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;type:BIGINT"`
		PackageId *int       `json:"package_id" gorm:"size:50;"`
		Name      *string    `json:"name" gorm:"size:50;"`
		Code      *string    `json:"code" gorm:"size:50;"`
		CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}

	SubscribeHistory struct {
		*gorm.DeletedAt

		Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;type:BIGINT"`
		UserId    *uuid.UUID `json:"user_id" gorm:"size:36;not null;"`
		PackageId *int       `json:"package_id" gorm:"size:50;"`
		ExpiredAt *time.Time `json:"expired_at" gorm:"type:timestamptz"`
		CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)
