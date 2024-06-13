package user

import (
	"github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	"github.com/google/uuid"
)

type (
	User struct {
		Id          uuid.UUID       `json:"id"`
		FirstName   string          `json:"first_name"`
		LastName    string          `json:"last_name"`
		UserCode    string          `json:"user_code"`
		Email       string          `json:"email"`
		Phone       string          `json:"phone"`
		Image       string          `json:"image"`
		Dob         string          `json:"dob"`
		Domicilie   string          `json:"domicilie"`
		Address     string          `json:"address"`
		Gender      user.UserGender `json:"gender"`
		PackageId   int             `json:"package_id"`
		FeatureCode string          `json:"feature_code"`
	}
)
type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	RegistrationRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Image     string `json:"image"`
		Dob       string `json:"dob"`
		Domicilie string `json:"domicilie"`
		Address   string `json:"address"`
		Gender    string `json:"gender"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}
)
