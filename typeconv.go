// +build OMIT
package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	x := int(5)
	fmt.Println(x + 2) // what happens if we put 2.0?

	y := 4.5
	fmt.Println(x + y) // what type is y?

	var p interface{}
	p = new(Point)
	fmt.Println(p)
}

// ENDMAIN OMIT
