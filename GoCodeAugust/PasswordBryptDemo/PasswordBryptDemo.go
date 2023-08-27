package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password1 := "Password@123"
	hs, err := bcrypt.GenerateFromPassword([]byte(password1), bcrypt.MinCost)
	if err != nil{
		log.Println(err)
	}
	log.Println([]byte(password1))
	log.Println(hs)

	password2 := "Password@123"
	err = bcrypt.CompareHashAndPassword(hs, []byte(password2))
	if err != nil{
		log.Println("Password doesn't match")
	} else{
		log.Println("Password matches")
	}
}
