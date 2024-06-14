package user

import (
	"context"
	"math/rand"
	"time"

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

// GetRandom implements UserRepositoryContract.
func (r *UserRepository) GetRandom(ctx context.Context, gender userEntity.UserGender, userCode string, execptionUserCode []string) (*userEntity.User, error) {
	var (
		count            int64
		user             userEntity.User
		errCount, errGet error
	)

	lenExeption := len(execptionUserCode)

	if lenExeption == 0 {
		errCount = r.db.Debug().Model(userEntity.User{}).Where("gender != ? AND user_code != ?", string(gender), userCode).Count(&count).Error
	} else {
		errCount = r.db.Debug().Model(userEntity.User{}).Where("gender != ? AND user_code != ? AND user_code NOT IN ?", string(gender), userCode, execptionUserCode).Count(&count).Error
	}

	if errCount != nil {
		return nil, errCount
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	offset := rng.Intn(int(count))

	if lenExeption == 0 {
		errGet = r.db.Debug().Offset(offset).Where("gender != ? AND user_code != ?", string(gender), userCode).First(&user).Error
	} else {
		errGet = r.db.Debug().Offset(offset).Where("gender != ? AND user_code != ? AND user_code NOT IN ?", string(gender), userCode, execptionUserCode).First(&user).Error
	}

	if errGet != nil {
		return nil, errGet
	}
	return &user, nil
}
