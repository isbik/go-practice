package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	Age      int
}

var users []User
var Admin = User{Name: "admin", Password: "password"}

func getUserList(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(&users)

	if err != nil {
		return
	}

	fmt.Println("test")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(users) == 0 {
		w.Write([]byte("[]"))
	} else {
		w.Write(response)
	}
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var userId int

	userId, err := strconv.Atoi(vars["id"])

	if err == nil {
		fmt.Printf("i=%d", err)
	}

	for _, user := range users {
		if user.Id == userId {
			response, err := json.Marshal(&user)

			if err != nil {
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}
	}

	w.WriteHeader(404)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.Id = rand.Int()
	users = append(users, user)

	w.WriteHeader(201)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	var filteredUsers []User

	userId, err := strconv.Atoi(vars["id"])

	if err == nil {
		fmt.Printf("i=%d", err)
	}

	for _, user := range users {
		if user.Id != userId {
			filteredUsers = append(filteredUsers, user)
		}
	}

	if len(filteredUsers) == len(users) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		users = filteredUsers
	}

	w.WriteHeader(201)
}

func AuthMiddleware(h http.Handler) http.Handler {
	print("tet")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name, password, ok := r.BasicAuth()

		if name != Admin.Name || password != Admin.Password || !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}

		h.ServeHTTP(w, r)
	})
}

func main() {

	router := mux.NewRouter()

	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.Use(AuthMiddleware)

	userRouter.HandleFunc("/", getUserList).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9]+}", getUserById).Methods(http.MethodGet)

	userRouter.HandleFunc("/", createUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id:[0-9]+}", deleteUser).Methods(http.MethodDelete)

	http.Handle("/", router)

	http.ListenAndServe(":4000", nil)
}
