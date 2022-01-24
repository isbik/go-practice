package app

import (
	"fmt"
	"main/internal/models"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func Start(db *gorm.DB) {
	findMirOldCardOwners(db)
}

func createUser(db *gorm.DB) {
	user := &models.User{
		Name:     "Kate",
		Age:      14,
		IsVerify: true,

		Cards: []models.Card{{
			Number: "0000123412341234",
			Type:   models.MASTERCARD,
		}},
	}

	result := db.Create(user)

	if result.Error != nil {
		panic(result.Error)
	}
}

func createUsers(db *gorm.DB) {
	var users []models.User

	for i := 0; i < 30; i++ {
		user := models.User{
			Name:     gofakeit.Name(),
			Age:      gofakeit.Number(18, 70),
			IsVerify: gofakeit.Bool(),
		}

		for i := 0; i < gofakeit.Number(1, 3); i++ {
			user.Cards = append(user.Cards, models.Card{
				Number: gofakeit.CreditCard().Number,
				Type:   models.CardType(gofakeit.RandomString([]string{"MASTERCARD", "MIR", "VISA"})),
			})
		}

		users = append(users, user)
	}
	db.Create(&users)
}

func findUsersHasMoreTwoCards(db *gorm.DB) {

	var result []models.User

	subQuey := db.Model(&models.Card{}).Select("user_id").Group("user_id").Having("COUNT(*) > 2")
	db.Where("id IN (?)", subQuey).Find(&result)

	fmt.Println(result)
}

func findUsersHasCardVisa(db *gorm.DB) {

	var result []models.User

	subQuery := db.Model(&models.Card{}).Distinct("user_id").Where("type='VISA'")
	db.Where("id IN (?)", subQuery).Find(&result)

	fmt.Println(result)
}

func findMirOldCardOwners(db *gorm.DB) {
	var result []models.Card

	subQuery := db.Model(&models.User{}).Distinct("user_id").Where("age > 50")
	db.Where("user_id IN (?) AND type='MIR'", subQuery).Find(&result)

	fmt.Println(result)
}
