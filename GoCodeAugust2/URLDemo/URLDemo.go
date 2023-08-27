package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var mp map[*url.URL]string
var chartset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789" 
func main(){
	myurl := "https://www.hello.example.com/home.html"
	result, err := url.Parse(myurl)
	if err != nil{
		log.Println(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	myurl2 := &url.URL{
		Scheme: "http",
		Host: "www.example.com",
		Path: "main.html",
	}
	fmt.Println(myurl2.String())

	xs := strings.Split(result.Host, ".")
	fmt.Println(xs)
	fmt.Println("Top Level Domain: ", xs[len(xs)-2]+"."+xs[len(xs)-1])

	rand.Seed(time.Now().Unix())

	mp = make(map[*url.URL]string)

	short := URLToShort(myurl2)
	fmt.Println(short)

	newurl := ShortToURL(short)
	fmt.Println(newurl)
}



func URLToShort(inputUrl *url.URL)string{
	short := make([]byte, 6)
	for i:= range short{
		short[i] = chartset[rand.Intn(len(chartset))]
	}
	hs,err := bcrypt.GenerateFromPassword(short, bcrypt.DefaultCost)
	if err != nil{
		log.Println(err)
	}
	mp[inputUrl] = string(hs)
	return string(short)
}


func ShortToURL(short string)*url.URL{
	for k := range mp{
		err := bcrypt.CompareHashAndPassword([]byte(mp[k]),[]byte(short))
		if err == nil{
			fmt.Println("Matches")
			return k
		}
	}
	return nil
}