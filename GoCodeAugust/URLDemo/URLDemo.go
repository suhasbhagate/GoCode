package main

import (
	"fmt"
	"log"
	"net/url"
)

// const myurl string ="https://lco.dev:3000/learn=reactjs/sjdhf"
//const myurl = "https://www.abcd.com:1000/home.html"
const myurl = "https://example.com?b=2&a=1"

func main() {
	result, err := url.Parse(myurl)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result.Query())
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	newUrl := &url.URL{
		Scheme: "https",
		Host:   "www.abcd.com:1000",
		Path:   "home.html",
	}
	newUrlString := newUrl.String()
	fmt.Println(newUrlString)
}
