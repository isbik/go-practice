package models

import "gorm.io/gorm"

type User struct {
	ID       int
	Name     string
	Age      int
	IsVerify bool `gorm:"column:trusted"`
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})

	if err != nil {
		return err
	}

	return nil
}
