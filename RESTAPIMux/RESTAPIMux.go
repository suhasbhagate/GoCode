package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type User struct{
	UserID int `json:userid`
	UserName string `json:"username`
	EmailID string `json:"emailid`
	Role string `json:"role`
	Status string `json:"status`
}

var users []User

func (c *User)IsEmpty() bool{
	var b interface{} = nil
	return c.UserID == b && c.UserName==""
}

func main(){
	fmt.Println("User Management API")
	r := mux.NewRouter()
	users = append(users, User{UserID: 1 , UserName: "Suhas", EmailID:"sbb@gmail.com",Role:"Technical Lead",Status:"Approved"})
	users = append(users, User{UserID: 2 , UserName: "Saksham", EmailID:"ssb@gmail.com",Role:"Senior Developer",Status:"Approved"})
	users = append(users, User{UserID: 3 , UserName: "Atharva", EmailID:"aab@gmail.com",Role:"Developer",Status:"Pending"})

	r.HandleFunc("/",ServeHome).Methods("GET")
	r.HandleFunc("/users",getAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}",getOneUser).Methods("GET")
	r.HandleFunc("/user",createOneUser).Methods("POST")
	r.HandleFunc("/user/{id}",updateOneUser).Methods("PUT")
	r.HandleFunc("/user/{id}",deleteOneUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000",r))

}

func ServeHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome to User Management API</h1>"))
}

func getAllUsers(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get All Users")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(users)
}

func getOneUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Get One User")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	//fmt.Printf("Type of params is: %T",params)
	currentUserID, _ := strconv.Atoi(params["id"])

	for _,user := range users{
		if user.UserID == currentUserID{
			json.NewEncoder(w).Encode(user)
			return
		} 
	}
	json.NewEncoder(w).Encode("No user found")
	return
}

func createOneUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Create One User")
	w.Header().Set("Content-Type","application/json")

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	if user.IsEmpty(){
		json.NewEncoder(w).Encode("Please send data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	user.UserID = rand.Intn(100)

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
	return
}

func updateOneUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Update One User")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	currentUserID, _ := strconv.Atoi(params["id"])

	for index,user := range users{
		if user.UserID == currentUserID{
			users = append(users[:index],users[index+1:]...)
			var user User
			_ = json.NewDecoder(r.Body).Decode(&user)
			user.UserID = currentUserID
			users = append(users, user)
			json.NewEncoder(w).Encode(user)
			return
		} 
	}
	json.NewEncoder(w).Encode("No user found")
	return
}

func deleteOneUser(w http.ResponseWriter , r *http.Request){
	fmt.Println("Delete One User")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	currentUserID, _ := strconv.Atoi(params["id"])

	for index,user := range users{
		if user.UserID == currentUserID{
			users = append(users[:index],users[index+1:]...)
			fmt.Fprintf(w, "User with UserID %v has been deleted.\n",currentUserID)
			//break
		} 
	}
}