package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"main/internal/middleware"
	"main/internal/mocks"
)

func getUserList(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(&mocks.Users)

	if err != nil {
		return
	}

	fmt.Println("test")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if len(mocks.Users) == 0 {
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

	for _, user := range mocks.Users {
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
	var user mocks.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if id := r.Context().Value("ID"); id != nil {
		if value, ok := id.(int); ok {
			user.Id = value
		}
	}

	mocks.Users = append(mocks.Users, user)

	w.WriteHeader(201)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	var filteredUsers []mocks.User

	userId, err := strconv.Atoi(vars["id"])

	if err == nil {
		fmt.Printf("i=%d", err)
	}

	for _, user := range mocks.Users {
		if user.Id != userId {
			filteredUsers = append(filteredUsers, user)
		}
	}

	if len(filteredUsers) == len(mocks.Users) {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		mocks.Users = filteredUsers
	}

	w.WriteHeader(201)
}

func main() {

	router := mux.NewRouter()

	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.Use(middleware.AuthMiddleware)
	userRouter.Use(middleware.RequestIDMiddleware)

	userRouter.HandleFunc("/", getUserList).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id:[0-9]+}", getUserById).Methods(http.MethodGet)

	userRouter.HandleFunc("/", createUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id:[0-9]+}", deleteUser).Methods(http.MethodDelete)

	http.Handle("/", router)

	http.ListenAndServe(":4000", nil)
}
