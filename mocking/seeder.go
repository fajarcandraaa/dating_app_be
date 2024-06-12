package mocking

import (
	"fmt"

	"gorm.io/gorm"
)

func seedUser(db *gorm.DB) error {
	fakeUserData := FakeUsers()
	for _, data := range fakeUserData {
		err := db.Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedAccount(db *gorm.DB) error {
	fakeAccountData := FakeAccountData()
	err := db.Create(&fakeAccountData).Error
	if err != nil {
		return err
	}

	return nil
}

func seedPackage(db *gorm.DB) error {
	fakePackages := FakePackageData()
	for _, data := range fakePackages {
		err := db.Debug().Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func seedPremiFeature(db *gorm.DB) error {
	fakePremiPackages := FakePremiumPackage()
	for _, data := range fakePremiPackages {
		err := db.Debug().Create(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func SeedData(db *gorm.DB) error {
	var (
		countUser int64
		countAccount int64
		countPackages int64
		countFeatures int64
	)

	go func() {
		db.Table("users").Count(&countUser)
		if countUser == 0 {
			err := seedUser(db)
			if err != nil {
				fmt.Printf("errors when seeding user table : %s", err.Error())
			}
	
			return
		}
	}()

	go func() {
		db.Table("accounts").Count(&countAccount)
		if countAccount == 0 {
			err := seedAccount(db)
			if err != nil {
				fmt.Printf("errors when seeding account table : %s", err.Error())
			}
	
			return
		}
	}()

	go func() {
		db.Table("packages").Count(&countPackages)
		if countPackages == 0 {
			err := seedPackage(db)
			if err != nil {
				fmt.Printf("errors when seeding packages table : %s", err.Error())
			}
	
			return
		}
	}()

	go func() {
		db.Table("premium_features").Count(&countFeatures)
		if countFeatures == 0 {
			err := seedPremiFeature(db)
			if err != nil {
				fmt.Printf("errors when seeding premium_features table : %s", err.Error())
			}
	
			return
		}
	}()

	return nil
}
