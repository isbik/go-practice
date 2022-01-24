package main

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	mongo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Name       string `json:"name" bson:"name"`
	Sex        int    `json:"sex" bson:"sex"`
	Smartphone struct {
		Model  string `json:"model" bson:"model"`
		Vendor int    `json:"vendor" bson:"vendor"`
	}
	DateOfBirth int64 `json:"date_of_birth" bson:"date_of_birth"`
}

func main() {

	sess, err := mongo.Dial("mongodb://localhost:27017/test")

	CreateUsers(sess)

	if err != nil {
		panic(err)
	}

}

func CreateUsers(sess *mongo.Session) {
	for i := 0; i < 30; i++ {
		user := User{
			Name:        gofakeit.Name(),
			Sex:         gofakeit.Number(0, 1),
			DateOfBirth: gofakeit.Date().Unix() - int64(time.Hour)*18,
		}

		user.Smartphone.Model = gofakeit.Word()
		user.Smartphone.Vendor = gofakeit.Number(10, 100)

		err := Store(sess, user)

		if err != nil {
			fmt.Println(err)
		}

	}

}

func Store(sess *mongo.Session, user User) error {
	err := sess.DB("").C("users").Insert(user)

	if err != nil {
		return err
	}
	return err
}

func Find(sess *mongo.Session, name string) (User, error) {
	var user User

	q := bson.M{
		"name": name,
	}

	err := sess.DB("").C("users").Find(q).One(&user)

	if err != nil {
		return user, nil
	}

	return user, nil
}
