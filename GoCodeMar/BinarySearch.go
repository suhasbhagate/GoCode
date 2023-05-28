package main

import(
	"fmt"
)


func main(){
	numss := []int{1,3,5,7,9,10}
	fmt.Println(BinarySearch(numss,0,len(numss)-1, 6))
	//fmt.Println(BinarySearch(slc,0,len(slc)-1, 9))
}


func BinarySearch(numss []int, st, end, keyy int)int{
	if st<=end{
	mid := (st+end)/2
	if numss[mid]==keyy{
		return mid
	} else if keyy<numss[mid]{
		return BinarySearch(numss,st,mid-1,keyy)
	} else if keyy>numss[mid]{
		return BinarySearch(numss,mid+1,end,keyy)
	}
	}
	return -1
}