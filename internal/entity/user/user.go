package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Error string

const (
	//User Error
	ErrUserNotExist             = Error("domain.user.error.not_exist")
	ErrUserAlreadyExist         = Error("domain.user.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist  = Error("domain.user.error.credential_not_exist")
	ErrUsersUnprocessableEntity = Error("domain.user.error.unprocessable_entity")
)

type (
	User struct {
		*gorm.DeletedAt

		Id          *uuid.UUID `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
		FirstName   *string    `json:"first_name" gorm:"size:255;"`
		LastName    *string    `json:"last_name" gorm:"size:255;"`
		UserCode    *string    `json:"user_code" gorm:"size:20;unique"`
		Email       *string    `json:"email" gorm:"size:255;"`
		Phone       *string    `json:"phone" gorm:"size:255;"`
		Image       *string    `json:"image" gorm:"size:255;"`
		Dob         *string    `json:"dob" gorm:"size:255;"`
		Domicilie   *string    `json:"domicilie" gorm:"size:255;"`
		Address     *string    `json:"address" gorm:"size:255;"`
		PackageId   *int       `json:"package_id" gorm:"size:10;"`
		FeatureCode *string    `json:"feature_code" gorm:"size:50;"`
		CreatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt   *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)
