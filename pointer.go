// +build OMIT

package main

import "fmt"

type Point struct {
	x float64
	y float64
}

func main() {
	p0 := Point{} // stack
	p1 := p0

	p1.x = 5
	fmt.Println(p0, p1)

	p3 := new(Point) // heap
	p4 := p3

	p3.x = 5
	fmt.Println(p3, p4)
}

// ENDMAIN OMIT
