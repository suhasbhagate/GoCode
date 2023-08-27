package main

import(
	"fmt"
)

func main(){
	slc := []int{1,2,3,4,5,1,2,3}
	mp := make(map[int]int)
	for _, v := range slc{
		mp[v]++
	}
	fmt.Println(mp)
	mp[4]--
	if mp[4]==0{
		delete(mp,4)
	}
	fmt.Println(mp)
	if _,ok := mp[5]; ok{
		delete(mp,5)
	}
	fmt.Println(mp)
}