package main

import(
	"fmt"
)

func BinarySearch(lst []int, start, end, key int) int{
	if start == end{
		if key == lst[start]{
			return start
		}
	} else{
		mid := (start + end)/2
		if key == lst[mid]{
			return mid
		} else {
			if key < lst[mid]{
				return BinarySearch(lst, start, mid-1,key)
			} else{
				return BinarySearch(lst, mid+1, end,key)
			}
		}
	}
	return -1
}

func main(){
	slc := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(BinarySearch(slc,0,8,5))
	fmt.Println(BinarySearch(slc,0,8,10))
}

