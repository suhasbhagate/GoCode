package main

import "fmt"

func main() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13
	m["k3"] = 17
    m["k4"] = 23

    fmt.Println("map:", m)

    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))

    delete(m, "k2")
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

	if val, flg := m["k2"]; flg{
		fmt.Println("Value %v Found", val)
	} else{
		fmt.Println("Value not found")
	}

	for k, v := range m{
		fmt.Printf("Index: %v Value: %v\n",k,v)
	}

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}