package dto

import (
	"fmt"
	"time"

	"github.com/fajarcandraaa/dating_app_be/helpers"
	accountEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/account"
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
	"github.com/google/uuid"
)

func TransformEntityToPresentation(user *userEntity.User) userPresentation.User {
	var resp userPresentation.User

	if user.Id != nil {
		resp.Id = *user.Id
	}
	if user.FirstName != nil {
		resp.FirstName = *user.FirstName
	}
	if user.LastName != nil {
		resp.LastName = *user.LastName
	}
	if user.UserCode != nil {
		resp.UserCode = *user.UserCode
	}
	if user.Email != nil {
		resp.Email = *user.Email
	}
	if user.Phone != nil {
		resp.Phone = *user.Phone
	}
	if user.Image != nil {
		resp.Image = *user.Image
	}
	if user.Dob != nil {
		resp.Dob = *user.Dob
	}
	if user.Domicilie != nil {
		resp.Domicilie = *user.Domicilie
	}
	if user.Address != nil {
		resp.Address = *user.Address
	}
	if user.PackageId != nil {
		resp.PackageId = *user.PackageId
	}
	if user.FeatureCode != nil {
		resp.FeatureCode = *user.FeatureCode
	}

	return resp
}

func TransformPresentationToEntity(user userPresentation.User) userEntity.User {
	var resp userEntity.User
	if user.Id != uuid.Nil {
		resp.Id = &user.Id
	}
	if user.FirstName != "" {
		resp.FirstName = &user.FirstName
	}
	if user.LastName != "" {
		resp.LastName = &user.LastName
	}
	if user.UserCode != "" {
		resp.UserCode = &user.UserCode
	}
	if user.Email != "" {
		resp.Email = &user.Email
	}
	if user.Phone != "" {
		resp.Phone = &user.Phone
	}
	if user.Image != "" {
		resp.Image = &user.Image
	}
	if user.Dob != "" {
		resp.Dob = &user.Dob
	}
	if user.Domicilie != "" {
		resp.Domicilie = &user.Domicilie
	}
	if user.Address != "" {
		resp.Address = &user.Address
	}
	if user.PackageId != 0 {
		resp.PackageId = &user.PackageId
	}
	if user.FeatureCode != "" {
		resp.FeatureCode = &user.FeatureCode
	}

	return resp
}

func TransforRegistrationToDatabase(data userPresentation.RegistrationRequest) (*userEntity.User, *accountEntity.Account, error) {
	// var userGender string
	userId := uuid.New()
	userCode := helpers.GenerateUserCode()
	defaultPackageId := 1
	if data.Gender != string(userEntity.UserMale) {
		if data.Gender != string(userEntity.UserFemale) {
			return nil, nil, fmt.Errorf("gender is unidentified")
		}
	}

	gend := data.Gender
	respUser := userEntity.User{
		Id:          &userId,
		FirstName:   &data.FirstName,
		LastName:    &data.LastName,
		UserCode:    &userCode,
		Email:       &data.Email,
		Phone:       &data.Phone,
		Image:       &data.Image,
		Dob:         &data.Dob,
		Domicilie:   &data.Domicilie,
		Address:     &data.Address,
		Gender:      &gend,
		PackageId:   &defaultPackageId,
		FeatureCode: nil,
		CreatedAt:   &time.Time{},
		UpdatedAt:   &time.Time{},
	}

	hashPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		return nil, nil, err
	}
	respAccount := accountEntity.Account{
		Email:     data.Email,
		Username:  data.Username,
		Password:  hashPassword,
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
	}

	return &respUser, &respAccount, nil
}
