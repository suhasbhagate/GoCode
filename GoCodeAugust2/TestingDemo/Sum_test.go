package main

import(
	"testing"
)

func Test_Sum(t *testing.T){
	r := Sum(1,2,3)
	if r!= 6{
		t.Error("Expected ", 6, " Got ", r)
	}
}