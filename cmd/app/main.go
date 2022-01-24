package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	AggregateUsers(sess)

	if err != nil {
		panic(err)
	}
}

func AggregateUsers(sess *mongo.Session) {

	var result []bson.M

	pipeline := []bson.M{
		{
			"$group": bson.M{
				// TODO fix There are  nested object
				"_id": bson.M{
					"sex": "$sex", "model": "$smartphone.model",
				},
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
	}

	err := sess.DB("").C("users").Pipe(pipeline).All(&result)

	if err != nil {
		fmt.Print(err)
	}

	file, _ := json.MarshalIndent(result, "", " ")

	_ = ioutil.WriteFile("data.json", file, 0644)

}

func CreateUsers(sess *mongo.Session) {
	for i := 0; i < 30; i++ {
		user := User{
			Name:        gofakeit.Name(),
			Sex:         gofakeit.Number(0, 1),
			DateOfBirth: gofakeit.Date().Unix() - int64(time.Hour)*18,
		}

		user.Smartphone.Model = gofakeit.RandomString([]string{"Apple", "Mui", "Redmi"})
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
