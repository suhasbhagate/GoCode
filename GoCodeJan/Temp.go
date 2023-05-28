package main
import (
	"fmt"
	"reflect"
	"sort"
)

func main(){
	arr1 := []int{1,1,2,3,4,0}
	arr2 := []int{1,0,2,1,4,3}
	fmt.Println(arr1==arr2)
	fmt.Println(reflect.DeepEqual(arr1,arr2))
	sort.Ints(arr1)
	sort.Ints(arr2)
	fmt.Println(reflect.DeepEqual(arr1,arr2))
}
