package main

import (
	"main/internal/app"
	"main/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=test port=5431 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = models.InitModels(db)

	app.Start(db)

	if err != nil {
		panic(err)
	}

}
