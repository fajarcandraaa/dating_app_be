package user

import (
	"context"

	"github.com/pkg/errors"

	"github.com/fajarcandraaa/dating_app_be/config/database"
	"github.com/fajarcandraaa/dating_app_be/helpers"
	"github.com/fajarcandraaa/dating_app_be/helpers/errorcodehandling"
	"github.com/fajarcandraaa/dating_app_be/internal/dto"
	accountEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/account"
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	"gorm.io/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ UserRepositoryContract = &UserRepository{}

// SignIn implements UserRepositoryContract.
func (r *UserRepository) SignIn(ctx context.Context, payload userPresentation.LoginRequest) (*userPresentation.User, error) {
	var (
		userData    = userEntity.User{}
		accountData = accountEntity.Account{}
	)

	err := r.db.Model(accountEntity.Account{}).Where("username = ?", payload.Username).Take(&accountData).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case database.ErrNoRowsFound:
			return nil, database.ErrNoRowsFound
		case database.ErrUniqueViolation:
			return nil, database.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to login from database")
		}
	}

	cekPass := helpers.CheckPasswordHash(payload.Password, accountData.Password)
	if !cekPass {
		return nil, errors.Errorf("password is wrong")
	}

	err = r.db.Model(userEntity.User{}).Where("email = ?", accountData.Email).Take(&userData).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case database.ErrNoRowsFound:
			return nil, database.ErrNoRowsFound
		case database.ErrUniqueViolation:
			return nil, database.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to get user data from database")
		}
	}

	dtoUser := dto.TransformEntityToPresentation(&userData)

	return &dtoUser, nil
}

// SignUp implements UserRepositoryContract.
func (r *UserRepository) SignUp(ctx context.Context, payload userEntity.User, payloadAccount accountEntity.Account) error {
	tx := r.db.Begin()
	err := tx.Create(&payload).Error
	if err != nil {
		tx.Rollback()
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case database.ErrUniqueViolation:
			return userEntity.ErrUserAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert user from database")
		}
	}

	err = tx.Create(&payloadAccount).Error
	if err != nil {
		tx.Rollback()
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case database.ErrUniqueViolation:
			return userEntity.ErrUserAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert user account from database")
		}
	}

	tx.Commit()
	return nil
}

// NewAccount implements UserRepositoryContract.
// func (r *UserRepository) newAccount(ctx context.Context, payload accountEntity.Account) error {
// 	err := r.db.Create(&payload).Error
// 	if err != nil {
// 		parsed := r.codeError.ParseSQLError(err)
// 		switch parsed {
// 		case database.ErrUniqueViolation:
// 			return userEntity.ErrUserAlreadyExist
// 		default:
// 			return errors.Wrap(parsed, "build statement query to insert user account from database")
// 		}
// 	}
// 	return nil
// }
