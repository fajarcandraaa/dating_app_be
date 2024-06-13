package dto

import (
	userEntity "github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	userPresentation "github.com/fajarcandraaa/dating_app_be/internal/presentation/user"
)

func UserEntityToPresent(user *userEntity.User) userPresentation.User {
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