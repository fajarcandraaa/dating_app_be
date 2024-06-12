package database

import (
	"github.com/fajarcandraaa/dating_app_be/internal/entity/account"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/subscription"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/user"
)

func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&user.User{},
		&account.Account{},
		&subscription.Package{},
		&subscription.PremiumFeature{},
		&subscription.SubscribeHistory{},
	}

	return migrationData
}
