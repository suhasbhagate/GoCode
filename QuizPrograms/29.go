package main
import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50}
	y := x[1:3]
	fmt.Println(y)
}