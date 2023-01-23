package main

import(
	"fmt"
	//"bytes"
	"sort"
	"reflect"
)

func main(){
	// a:=[]byte{1,2,3}
	// b:=[]byte{3,2,1}
	// fmt.Printf("Slice a and b are Same %v\n",bytes.Compare(a,b))

	a := []int{1,2,3}
	b := []int{3,2,1}
	sort.Ints(a)
	sort.Ints(b)
	fmt.Printf("Slice a and b are same %v\n",reflect.DeepEqual(a,b))

	type interface1 interface{
		Display();
	}

	type interface2 interface{
		//Display();
		//Draw();
	}

	var i1 interface1;
	var i2 interface2;
	fmt.Printf("Interface 1 and 2 are same %v\n",i1==i2)
	fmt.Printf("Interface 1 and 2 are same %v\n", reflect.DeepEqual(i1, i2))

	type struct1 struct{
		RollNo int
		Name string
		//marks []int
	} 

	st1 := struct1{
		RollNo: 10,
		Name:"Saksham",
	}
	
	st2 := struct1{
		RollNo: 10,
		Name: "Saksham",
	}

	fmt.Printf("Structure 1 and 2 are same %v\n", st1==st2)
	fmt.Printf("Structure 1 and 2 are same %v\n", reflect.DeepEqual(st1, st2))

	m1 := map[string]int{"One":1, "Two":2, "Three":3}
	m2 := map[string]int{"One":1, "Two":2, "Three":4}
	fmt.Printf("Maps m1 and m2 are same %v\n",MapCompare(m1,m2))

	ar1 :=[5]int{1,2,3,4,5}
	ar2 :=[5]int{1,2,5,4,3}
	fmt.Printf("Array ar1 and ar2 are same %v\n",ar1==ar2)

}


func MapCompare(m1, m2 map[string]int)bool{
	if len(m1) != len(m2){
		return false
	}
	for k1, v1:= range m1{
		if m2[k1] != v1{
			return false
		}
	}
	return true
}