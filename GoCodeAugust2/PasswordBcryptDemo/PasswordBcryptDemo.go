package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main(){
	rand.Seed(time.Now().Unix())
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	pass := make([]byte, 8)
	for i := range pass{
		pass[i] = charset[rand.Intn(len(charset))]
	}
	fmt.Println(string(pass))
	hs, err := bcrypt.GenerateFromPassword(pass,bcrypt.DefaultCost)
	if err!=nil{
		log.Println(err)
	}
	
	err = bcrypt.CompareHashAndPassword(hs,pass)
	if err!=nil{
		log.Println(err)
	} else{
		fmt.Println("Password Matches")
	}
}