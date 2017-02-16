// +build OMIT

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Println("Element:")
	for x := range s {
		fmt.Println(x)
	}

	fmt.Println("Index and element:")
	for i, x := range s {
		fmt.Println(i, x)
	}

	m := map[string]int{
		"foo": 6,
		"bar": 7,
	}
	fmt.Println("Key and value:")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// ENDMAIN OMIT
