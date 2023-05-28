package main

import(
	"fmt"
)

func main(){
	// f := 10.67
	// fmt.Printf("Value is %.6f",f)
	arr := [...]int{-4, 3, -9, 0, 4, 1}

	poscnt := 0
    negcnt := 0
    zerocnt := 0
    for _,v:= range arr{
        if v > 0{
            poscnt++
        } else if v < 0 {
            negcnt++
        } else {
            zerocnt++
        }
    }
	posper := float32(poscnt)/float32(len(arr))
	negper := float32(negcnt)/float32(len(arr))
	zeroper := float32(zerocnt)/float32(len(arr))

    // fmt.Printf("%.6f",float32(poscnt/len(arr)))
    // fmt.Printf("%.6f",float32(negcnt/len(arr)))
    // fmt.Printf("%.6f",float32(zerocnt/len(arr)))
	fmt.Printf("%.6f",posper)
	fmt.Printf("%.6f",negper)
	fmt.Printf("%.6f",zeroper)
}