// +build OMIT

package main

import "fmt"

func main() {
	// array
	a1 := [3]int{}
	a2 := [...]int{1, 2, 3}
	fmt.Println(a1, a2)

	// string
	u1 := "Chrząszcz"
	r1 := []rune(u1)
	fmt.Println(len(u1), len(r1))

	// slice
	s1 := a2[1:]
	s2 := []int{1, 2, 3}
	s3 := make([]int, 3, 10)
	fmt.Println(s1, s2, s3)
	fmt.Println(len(s3), cap(s3))
}

// ENDMAIN OMIT
