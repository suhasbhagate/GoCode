package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	str := "12:05:45AM"
	var outputtime string = ""
	amtime, _, fndam := strings.Cut(str,"AM")
	if fndam{
		slc := strings.Split(amtime,":")
		if slc[0]=="12"{
			outputtime = "00"
		} else {
			outputtime = slc[0]
		}
		outputtime = outputtime+":"+slc[1]+":"+slc[2]
		fmt.Println(outputtime)
	}

	pmtime, _, fndpm := strings.Cut(str,"PM")
	if fndpm{
		outputtime = ""
		slc := strings.Split(pmtime,":")
		addhr := 0
		hrs,err := strconv.Atoi(slc[0])
		if err!= nil{
			fmt.Errorf("Cannot convert string to int")
		}

		if hrs == 12{
			addhr = 0
		} else {
			addhr = 12
		}
		hrs += addhr
		outputtime = strconv.Itoa(hrs)
		outputtime = outputtime + ":"+slc[1]+":"+slc[2]
		fmt.Println(outputtime)
	}

}