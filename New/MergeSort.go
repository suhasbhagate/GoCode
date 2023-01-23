package main

import(
	"fmt"
)

var ans []int
func MergeSort(lst []int, low, high int){
	if low < high{
		mid := (low + high) / 2
		MergeSort(lst, low, mid)
		MergeSort(lst,mid+1, high)
		Merge(lst, low, mid, high) 
	}
}

func Merge(lst []int, low, mid, high int){
	i,j := low, high
	//ans := make([]int, len(lst))

	for i<= mid && j <=high{
		if lst[i] < lst[j]{
			ans = append(ans, lst[i])
			i++
		} else{
			ans = append(ans, lst[j])
			j++
		}
	}
	if i>mid{
		for k:=j; k<high;k++{
			ans = append(ans, lst[k])
		}
	} else {
		for k:= i; k<mid; k++{
			ans = append(ans, lst[i])
		}
	}
	//fmt.Println(ans)
}

func main(){
	slc := []int{4,6,2,6,8,3,9}
	MergeSort(slc,0,6)
	fmt.Println(ans)
}