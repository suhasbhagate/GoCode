package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId   int
	UserName string
}

var userList = []User{{11, "Saksham"}, {12, "Suhas"}}

func main() {
	r := gin.Default()
	r.GET("/getuser/:id", getuser)
	r.POST("/adduser", adduser)
	r.Run("localhost:2400")
}

func getuser(g *gin.Context) {
	id,_ := strconv.Atoi(g.Param("id"))
	for _, v := range userList{
		if id== v.UserId{
			g.IndentedJSON(http.StatusFound,v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound,gin.H{"message":"Element Not Found"})
}

func adduser(g *gin.Context) {
	var usr User
	err := g.BindJSON(&usr)
	if err != nil {
		log.Println(err)
	}
	userList = append(userList, usr)
	g.IndentedJSON(http.StatusOK, userList)
}
