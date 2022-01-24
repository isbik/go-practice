package models

import (
	"database/sql/driver"

	"gorm.io/gorm"
)

type CardType string

const (
	VISA       CardType = "VISA"
	MASTERCARD CardType = "MASTERCARD"
	MIR        CardType = "MIR"
)

func (ct CardType) Value() (driver.Value, error) {
	return string(ct), nil
}

type User struct {
	ID       int
	Name     string
	Age      int
	IsVerify bool   `gorm:"column:trusted"`
	Cards    []Card `gorm:"foreignKey:UserID"`
}

type Card struct {
	ID     int
	Number string
	Type   CardType `sql:"card_type"`
	UserID int
}

func InitModels(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Card{})

	if err != nil {
		return err
	}

	return nil
}
