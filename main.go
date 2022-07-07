package main

import (
	"fmt"
	"log"
	"net/http"
)

const hashCost = 8

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("testing"))
	fmt.Println(w)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin bitch")
}
