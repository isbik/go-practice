package main

import (
	"fmt"
	"net/http"
	"time"
)

type ResponseHandler struct {
	message string
}

type ResponseTimeoutHandler struct{}

// Send message
func (ch ResponseHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, ch.message)
}

// Timeout 2 seconds
func (rt ResponseTimeoutHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(2 * time.Second)
}

func main() {

	http.Handle("/hello", ResponseHandler{message: "hello"})

	http.Handle("/source", http.FileServer(http.Dir("../")))

	longPill := ResponseTimeoutHandler{}

	http.Handle("/longPill", http.TimeoutHandler(longPill, time.Second, "Request timeout"))

	http.Handle("/", http.RedirectHandler("https://google.com", http.StatusMovedPermanently))

	http.ListenAndServe(":4000", nil)
}
