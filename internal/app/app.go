package app

import (
	"fmt"
	"main/internal/models"

	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	selectAndUpdate(db)
}

func create(db *gorm.DB) {
	user := &models.User{
		Name:     "test",
		Age:      12,
		IsVerify: true,
	}

	result := db.Create(user)

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("User created with id %d", user.ID)
}

func selectAndUpdate(db *gorm.DB) {
	user := &models.User{}

	db.Where("name = ?", "test").First(&user)

	if user.ID > 0 {
		user.Name = "Dmitry"
		db.Save(user)
	}

	fmt.Printf("User updated with id %d", user.ID)

	// or this way but all records
	// db.Model(&models.User{}).Where("name = ?", "test").Update("name", "Dmitry")

}
