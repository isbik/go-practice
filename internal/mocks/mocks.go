package mocks

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Age      int
}

var Users []User
var Admin = User{Name: "admin", Password: "password"}
