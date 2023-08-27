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

var userList = []User{{1, "Saksham"}, {2, "Suhas"}}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", ServerHome).Methods("GET")
	r.HandleFunc("/adduser", adduser).Methods("POST")
	r.HandleFunc("/getuser/{id}", getuser).Methods("GET")
	log.Fatal(http.ListenAndServe("localhost:2300", r))
}

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
}

func adduser(w http.ResponseWriter, r *http.Request) {
	var usr User
	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		log.Println(err)
	}
	userList = append(userList, usr)
	json.NewEncoder(w).Encode(userList)
}

func getuser(w http.ResponseWriter, r *http.Request) {
	Param := mux.Vars(r)
	id, _ := strconv.Atoi(Param["id"])
	for _, v := range userList {
		if id == v.UserId {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	json.NewEncoder(w).Encode("Element No Found")
}
