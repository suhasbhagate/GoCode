package main

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)
var people []Person
type Person struct{
	Id 		int
	Name 	string
	Age 	int
}

func main(){
	people = []Person{{Id:10, Name:"Saksham", Age:15,},{Id:15,Name:"Suhas",Age:40,},{Id:20,Name:"Vihan",Age:10}}
	router := gin.Default()
	router.GET("/getallperson",getAllPerson)
	router.POST("/addperson", AddPerson)
	router.GET("/getperson/:Id", GetPerson)
	router.Run(":8181")
}

func getAllPerson(c *gin.Context){
	c.IndentedJSON(http.StatusOK, people)
}

func AddPerson(c *gin.Context){
	var newperson Person
	err := c.BindJSON(&newperson)
	if err!=nil{
		return
	}
	people = append(people,newperson)
	c.IndentedJSON(http.StatusCreated, newperson)
}

func GetPerson(c *gin.Context){
	id := c.Param("Id")
	intit, _ :=strconv.Atoi(id) 
	for _,a := range people{
		if a.Id== intit{
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Person not found"})
}