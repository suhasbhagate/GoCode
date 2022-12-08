package main
import (
    "sync"
	"fmt"
)
var once sync.Once

type single struct {
	EmpID	int
	EmpName	string
	Salary	float64
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        once.Do(
            func() {
                fmt.Println("Creting Single Instance Now")
                singleInstance = &single{}
            })
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}