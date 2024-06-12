package account

import (
	"time"

	"gorm.io/gorm"
)

type (
	Account struct {
		*gorm.DeletedAt

		Id        int64      `json:"id" gorm:"primaryKey;autoIncrement;type:BIGINT"`
		Email     string     `json:"email" gorm:"size:255;"`
		Username  string     `json:"username" gorm:"size:255;"`
		Password  string     `json:"password" gorm:"size:255;"`
		CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
		UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	}
)
