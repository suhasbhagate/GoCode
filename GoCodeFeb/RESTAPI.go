package main

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Firstname string
	Lastname  string
	Id        int
}

var people []Person

func main() {
	people = []Person{{"Suhas", "Bhagate", 41}, {"Saksham", "Bhagate", 14}}
	router := gin.Default()
	router.GET("/getallperson", getallperson)
	router.GET("/getperson/:id", getperson)
	router.POST("/addperson", addperson)
	router.Run("localhost:8282")
}

func getallperson(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, people)
}

func getperson(g *gin.Context) {
	i := g.Param("Id")
	id, err := strconv.Atoi(i)
	if err != nil {
		return
	}
	for _, v := range people {
		if v.Id == id {
			g.IndentedJSON(http.StatusOK, v)
		}
	}
	g.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person Not Found"})
}

func addperson(g *gin.Context) {
	var newperson Person
	err := g.BindJSON(&newperson)
	if err != nil {
		return
	}
	people = append(people, newperson)
	g.IndentedJSON(http.StatusCreated, newperson)
}
