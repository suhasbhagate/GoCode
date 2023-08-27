package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	UserId   int
	UserName string
}

var UserList []User = []User{{10, "Saksham"}, {20, "Suhas"}}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/getusers", getAllUsers).Methods("GET")
	r.HandleFunc("/getuser/{id}", getUserById).Methods("GET")
	r.HandleFunc("/adduser", addUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":4100", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home Page</h1>"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(UserList)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.Atoi(params["id"])
	for _, v := range UserList {
		if v.UserId == userId {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("No user found")
	return
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		return
	}
	UserList = append(UserList, usr)
	json.NewEncoder(w).Encode(UserList)
}
