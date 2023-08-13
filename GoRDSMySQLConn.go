package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Emp struct{
	EmpId		int32
	FirstName	string
	LastName	string
}

func main(){
	fmt.Println("Go MySQL Connection")

	db, err := sql.Open("mysql","admin:admin123@tcp(sbb-db1.cwnnghp1iejy.us-east-1.rds.amazonaws.com:3306)/sbbschema")

	if err!=nil{
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Successfully Connected to MySQL Database")

	/*
	insert, err := db.Query("INSERT INTO sbbschema.Employee values(30,'Atharva','Bhagate')")

	if err !=nil{
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("Successfully Inserted into Table")
	*/
	eid := 10
	fnm := "Suhas"
	res, err := db.Query("SELECT * FROM sbbschema.Employee WHERE EmpId=? and FirstName=?",eid,fnm)
	if err !=nil{
		panic(err.Error())
	}

	for res.Next(){
		var emp Emp
		err = res.Scan(&emp.EmpId,&emp.FirstName,&emp.LastName)
		if err !=nil{
			panic(err.Error())
		}
		fmt.Println(emp)
	}

}