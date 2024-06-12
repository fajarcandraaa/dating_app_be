package mocking

import (
	"fmt"

	"github.com/fajarcandraaa/dating_app_be/helpers"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/account"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/subscription"
	"github.com/fajarcandraaa/dating_app_be/internal/entity/user"
	"github.com/google/uuid"
)

var (
	FakeFirstName        string = "User"
	FakePhone            string = "0812345678910"
	FakeDob              string = "01/01/1980"
	FakeDomicilie        string = "Indonesia"
	FakeAddress          string = "Jl.Raya Hati-Hati Ya"
	FakePackageId        int    = 1
	FakePremiumPackageId int    = 2
	FakePacakgeName01    string = "Standart Package"
	FakePacakgeName02    string = "Premium Package"
	FakeFeatureName01    string = "No Quota"
	FakeFeatureCode01    string = "NQA"
	FakeFeatureName02    string = "Verified Label"
	FakeFeatureCode02    string = "VL"
	FakeAccountPassowrd  string = "123456"
)

// FakeUsers is used to create mocking data for table users
func FakeUsers() []*user.User {
	var (
		fakeUsersData []*user.User
	)

	for i := 0; i <= 10; i++ {
		fakeUserId := uuid.New()
		fakeLastName := fmt.Sprintf("00%d", i)
		fakeUserCode := helpers.GenerateUserCode()
		fakeEmail := fmt.Sprintf("user%s@email.com", fakeLastName)

		mockDataUsers := &user.User{
			Id:        &fakeUserId,
			FirstName: &FakeFirstName,
			LastName:  &fakeLastName,
			UserCode:  &fakeUserCode,
			Email:     &fakeEmail,
			Phone:     &FakePhone,
			Dob:       &FakeDob,
			Domicilie: &FakeDomicilie,
			Address:   &FakeAddress,
			PackageId: &FakePackageId,
		}

		fakeUsersData = append(fakeUsersData, mockDataUsers)
	}

	return fakeUsersData
}

func FakeAccountData() *account.Account {
	psswd, err := helpers.HashPassword(FakeAccountPassowrd)
	if err != nil {
		return nil
	}
	fakeAccount := &account.Account{
		Email:    "user001@email.com",
		Username: "user001",
		Password: psswd,
	}

	return fakeAccount
}

// FakePackageData is used to create mocking data for table packages
func FakePackageData() []*subscription.Package {
	var (
		fakePackageData []*subscription.Package
	)

	mockDataPackages01 := &subscription.Package{
		Name: &FakePacakgeName01,
	}
	fakePackageData = append(fakePackageData, mockDataPackages01)

	mockDataPackages02 := &subscription.Package{
		Name: &FakePacakgeName02,
	}
	fakePackageData = append(fakePackageData, mockDataPackages02)

	return fakePackageData
}

// FakePremiumPackage is used to create mocking data for table premium_features
func FakePremiumPackage() []*subscription.PremiumFeature {
	var (
		fakePremiumPackages []*subscription.PremiumFeature
	)

	mockPremiData01 := &subscription.PremiumFeature{
		PackageId: &FakePremiumPackageId,
		Name:      &FakeFeatureName01,
		Code:      &FakeFeatureCode01,
	}
	fakePremiumPackages = append(fakePremiumPackages, mockPremiData01)

	mockPremiData02 := &subscription.PremiumFeature{
		PackageId: &FakePremiumPackageId,
		Name:      &FakeFeatureName02,
		Code:      &FakeFeatureCode02,
	}
	fakePremiumPackages = append(fakePremiumPackages, mockPremiData02)

	return fakePremiumPackages
}
