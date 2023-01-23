package main

import (
	"strconv"
	pb "github.com/suhasbhagate/GoCode/Employee/proto"
)

type Emp struct {
	EmpId       int32
	FirstName   string
	LastName    string
	DisplayName string
	Age         int32
	Salary      float64
	Designation string
	Department  string
}

//Used for Search using Numeric Fields like EmpId, Age, Salary represented in form of string
type Emp1 struct {
	EmpId       string
	FirstName   string
	LastName    string
	DisplayName string
	Age         string
	Salary      string
	Designation string
	Department  string
}

func documentToEmployee(data *Emp) *pb.Employee {
	return &pb.Employee{
		EmpId:       data.EmpId,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		DisplayName: data.DisplayName,
		Age:         data.Age,
		Salary:      data.Salary,
		Designation: data.Designation,
		Department:  data.Department,
	}
}

//Used for Search using Numeric Fields like EmpId, Age, Salary represented in form of string
func documentToEmployee1(data *Emp1) *pb.Employee {
	eid,_ := strconv.Atoi(data.EmpId)
	age,_ := strconv.Atoi(data.Age)
	sal,_ := strconv.Atoi(data.Salary)
	return &pb.Employee{
		EmpId:       int32(eid),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		DisplayName: data.DisplayName,
		Age:         int32(age),
		Salary:      float64(sal),
		Designation: data.Designation,
		Department:  data.Department,
	}
}