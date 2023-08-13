package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)


type user struct{
	FirstName string
	MiddleName string
	LastName string
	EmailId string
	Age int
}

func main() {
	var userList []user
	f, err := excelize.OpenFile("D:\\GoCode\\ReadExcel\\SBB_Sample_Excel.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows, err := f.GetRows("Sample")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i:=1; i<len(rows); i++ {
		ageint,_ := strconv.Atoi(rows[i][4])
		usr :=user{
			FirstName: rows[i][0],
			MiddleName: rows[i][1],
			LastName: rows[i][2],
			EmailId: rows[i][3],
			Age: ageint,
		}
		userList = append(userList, usr)
	}
	fmt.Println(userList)
}