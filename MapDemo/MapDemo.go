package main

import (
	"fmt"
)

func main() {
	// mp := make(map[int]bool)
	// mp[1] = true
	// mp[2] = true
	// fmt.Println(mp)

	// if _,ok:=mp[2]; ok{
	// 	delete(mp,2)
	// } else{
	// 	fmt.Println("Not found")
	// }
	// fmt.Println(mp)

	mp := make(map[int]struct{})
	mp[1] = struct{}{}
	mp[2] = struct{}{}
	mp[3] = struct{}{}
	fmt.Println(mp)
	if _, ok := mp[3]; ok{
		fmt.Println("3 Exist in Map")
	}
	if _, ok := mp[4]; !ok{
		fmt.Println("4 Doesn't exist in Map")
	}
	delete (mp, 3)
	fmt.Println(mp)
}