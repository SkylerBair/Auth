package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
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

var usersMap map[string][]byte = make(map[string][]byte)

func Signup(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}                      //sets creds to a pointer to Credintals struct
	err := json.NewDecoder(r.Body).Decode(creds) //this decodes the jason request
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)

	usersMap[creds.Username] = hashedPassword

	// create a empty map to store the usernames and passwords in using the username as the key and the hashed password as the value.
	w.Write(hashedPassword)
	fmt.Println(w)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds) //
	if err != nil {
		log.Printf("Json Decoder Error %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//ugly to many ifs needs to be refactored into smaller functions
	if hashed, ok := usersMap[creds.Username]; ok {
		// compare hash and
		if err = bcrypt.CompareHashAndPassword(hashed, []byte(creds.Password)); err != nil {
			log.Printf("bcrypt Hash Error %v", err)

			w.WriteHeader(http.StatusUnauthorized)
			return
		} else {
			now := time.Now()
			token := bytes.NewBuffer([]byte(now.String()))
			h := sha256.New()
			_, err := h.Write(token.Bytes())
			if err != nil {
				log.Printf("Token return Error %v", err)

				w.WriteHeader(http.StatusBadGateway)
				return
			}
			b := h.Sum(nil)
			w.Write(b)
			return
		}

	} else {
		log.Printf("Bad Request Error %v", err)

		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
