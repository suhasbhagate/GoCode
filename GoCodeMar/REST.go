package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct{
	Firstname string
	Lastname string
	Id int
}

var people []Person = []Person{{"Saksham","Bhagate",11},{"Suhas","Bhagate",20}}

func main(){
	router := gin.Default()
	router.GET("/getallperson",getallperson)
	router.GET("/getperson/:id",getperson)
	router.POST("/addperson",addperson)
	router.Run("localhost:8989")
}

func getallperson(g *gin.Context){
	g.IndentedJSON(http.StatusOK,people)
}


func getperson(g *gin.Context){
	ids:=g.Param("id")
	idi,err := strconv.Atoi(ids)
	if err!=nil{
		return
	}
	for _,v:= range people{
		if idi == v.Id{
			g.IndentedJSON(http.StatusOK,v)
			return
		}
	}
	g.IndentedJSON(http.StatusNotFound,gin.H{"message":"Element not found"})
}


func addperson(g *gin.Context){
	var newp Person
	err := g.BindJSON(&newp)
	if err!=nil{
		return
	}
	g.IndentedJSON(http.StatusCreated,newp)
}