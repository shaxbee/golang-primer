// +build OMIT

package main

import "fmt"

func main() {
	// map
	m1 := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	m2 := m1
	m1["baz"] = 3
	fmt.Println(m1, m2)

	// empty map
	m3 := make(map[string]int, 100)
	fmt.Println(len(m3))

	// modeling set
	s1 := map[string]struct{}{
		"foo": {},
		"bar": {},
	}
	fmt.Println(s1)
}

// ENDMAIN OMIT
