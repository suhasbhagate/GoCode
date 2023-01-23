package main
import "fmt"

func main() {
	x := []int{1, 2, 3}
	for i := range x {
		x = append(x, i)
	}
	fmt.Println(x)
}