package app

import (
	"fmt"
	"main/internal/models"

	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	create(db)
}

func create(db *gorm.DB) {
	u := &models.User{
		Name: "test",
		Age:12,
		IsVerify: true,
	}

	result := db.Create(u)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("User created with id %d", u.ID)
}
