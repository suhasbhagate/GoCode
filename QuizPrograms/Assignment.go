package main
import (
        "fmt"
)

func main() {
        var x int
        arr := [3]int{3, 5, 2}
        x, arr = arr[0], arr[1:]

		x = arr[0]
		arr = arr[1:]
        fmt.Println(x, arr)
}
