package main

import(
	"fmt"
)



func main(){
	var mp [][]bool
	slc := []int{5, 10, 12, 13, 15, 18}
	x := make([]bool,len(slc))
	sm := 0
	for _,v:= range slc{
		sm +=v
	}
	fmt.Println("Sum: ", sm)
	SumOfSubset(slc, x, 0, 0, sm, 30, &mp)
	fmt.Println("Final Result",mp)
}

func SumOfSubset(w []int, x []bool, s int, k int, r int, m int, mp *[][]bool){	
	x[k] = true
	if ((s + w[k]) == m){
		for i:= k+1; i<len(w);i++{
			x[i] = false
		}
		res := make([]bool,len(w))
		copy(res,x)
		//fmt.Println(res)
		*mp = append(*mp, res)
	} else if ((s+w[k]+w[k+1])<= m){
		SumOfSubset(w, x, s+w[k], k+1, r-w[k], m,mp)
	} 
	if ((s + r - w[k] )>= m) && ((s+w[k+1]) <= m){
		x[k] = false
		SumOfSubset(w, x, s, k+1, r-w[k], m,mp)
	}
}
