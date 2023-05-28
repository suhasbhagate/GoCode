package main

import(
	"fmt"
	"sort"
)

func main(){
	slc:=[]int{5,3,6,9,2,7,1,2,8}
	// sort.Ints(slc)
	// fmt.Println(slc)
	sort.Sort(sort.IntSlice(slc))
	fmt.Println(slc)

	sort.Sort(sort.Reverse(sort.IntSlice(slc)))
	fmt.Println(slc)
}
