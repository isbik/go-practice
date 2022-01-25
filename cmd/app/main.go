package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Name     string
	Password string
}

type Wallet struct {
	Count int
	Id    int64
	Owner int
}

var Admin = User{Name: "admin", Password: "password"}

var Wallets = []Wallet{
	{
		Count: 1,
		Id:    1,
		Owner: 1,
	},
	{
		Count: 100,
		Id:    2,
		Owner: 2,
	},
	{
		Count: -100,
		Id:    3,
		Owner: 2,
	},
}

func main() {

	// You need query like this to get json data
	// curl 'localhost:4000/wallets?ids[]=1' --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ='
	http.HandleFunc("/wallets", func(rw http.ResponseWriter, r *http.Request) {

		name, pass, ok := r.BasicAuth()

		if name != Admin.Name || pass != Admin.Password || !ok {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
		}

		walletIds, ok := r.URL.Query()["ids[]"]

		if !ok {
			http.Error(rw, "invalid json", http.StatusBadRequest)
		}

		var result []Wallet

		for _, id := range walletIds {
			qId, err := strconv.ParseInt(id, 10, 12)

			if err != nil {
				panic(err)
			}

			for _, wallet := range Wallets {
				if wallet.Id == qId {
					result = append(result, wallet)
				}
			}
		}

		jsonData, err := json.Marshal(result)

		if err != nil {
			http.Error(rw, "invalid json", http.StatusInternalServerError)
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(jsonData)
	})

	http.ListenAndServe(":4000", nil)
}
