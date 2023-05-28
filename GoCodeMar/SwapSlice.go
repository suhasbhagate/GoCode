package main

import(
	"fmt"
)


func main(){
	slc:= []int{1,2,3,4,5,6}
	fmt.Println(Swap(slc))
}


func Swap(slc []int)[]int{
	for i,j:=0,len(slc)-1;i<j;i,j=i+1,j-1{
		slc[i],slc[j] = slc[j],slc[i]
	}
	return slc
}