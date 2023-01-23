package main
import (
"fmt"
)

type Writer interface {
	Write(string)
}
type StringWriter func(string) string

func (s StringWriter) Write(str string) { fmt.Println(s(str))
}

func main() {
	xyz := func(str string) string {
		return "XYZ-Stringer: " + str
	}
	var w Writer
	w = StringWriter(xyz)
	w.Write("Stringer")
}