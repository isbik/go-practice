package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World!")
	})
	fmt.Println("Test")
	http.ListenAndServe(":8080", nil)
}
