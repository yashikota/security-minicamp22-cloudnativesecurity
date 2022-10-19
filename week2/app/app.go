package main

import (
	"fmt"
	"net/http"
)

func getGreeting() string {
	return "Hello, security minicamp!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getGreeting())
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
