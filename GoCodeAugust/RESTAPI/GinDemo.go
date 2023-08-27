package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId   int
	UserName string
}

var UserList []User = []User{{10, "Saksham"}, {20, "Suhas"}}

func main() {
	r := gin.Default()
	r.GET("/getusers", getAllUsersGin)
	r.GET("/getuser/:id", getUserByIdGin)
	r.POST("/adduser", addUserGin)
	r.Run("localhost:4200")
}

func getAllUsersGin(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, UserList)
}

func getUserByIdGin(g *gin.Context) {
	id, _ := strconv.Atoi(g.Param("id"))
	for _, v := range UserList {
		if v.UserId == id {
			g.IndentedJSON(http.StatusFound, v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Element not found"})
}

func addUserGin(g *gin.Context) {
	var ur User
	err := g.BindJSON(&ur)
	if err != nil {
		return
	}
	UserList = append(UserList, ur)
	g.IndentedJSON(http.StatusOK, UserList)
}
