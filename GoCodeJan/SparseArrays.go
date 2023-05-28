package main

import(
	"fmt"
)

func main(){
	// strings := []string{"aba", "baba","aba","xzxb"}
	// queries := []string{"aba", "xzxb", "ab"}

	strings := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"}
	queries := []string{"abcde", "sdaklfj", "asdjf", "na", "basdn"}
	var arr []int32
    m := make(map[string]int)
    for _,v:= range queries{
        for _,u := range strings{
            if v==u{
                m[v]++
            }
        }
    }
    for _, k := range queries{
        arr = append(arr, int32(m[k]))
    }
    fmt.Println(arr)
}