package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var mp map[*url.URL]string
var charset string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	rand.Seed(time.Now().Unix())
	mp = make(map[*url.URL]string)
	myURL := &url.URL{
		Scheme: "https",
		Host:   "www.example.com",
		Path:   "home.html",
	}
	shorturl, err := URLToShort(myURL)
	if err != nil {
		fmt.Println("No Short URL Generated")
	} else {
		fmt.Println("Short URL: ", shorturl)
	}

	retURL, err := ShortToURL(shorturl)
	if err != nil {
		fmt.Println("No URL Generated from Short URL")
	} else {
		fmt.Println("URL: ", retURL)
	}
}

func URLToShort(longURL *url.URL) (string, error) {
	b := make([]byte, 6)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	enCryptedString, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while encrypting short url")
	}
	mp[longURL] = string(enCryptedString)
	return string(b), nil
}

func ShortToURL(short string) (*url.URL, error) {
	for k := range mp {
		err := bcrypt.CompareHashAndPassword([]byte(mp[k]), []byte(short))
		if err == nil {
			return k, nil
		}
	}
	return &url.URL{}, fmt.Errorf("URL not found")
}
