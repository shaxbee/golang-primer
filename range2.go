// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("Index and rune:")
	for i, c := range "Chrząszcz" {
		fmt.Printf("%d %#U\n", i, c)
	}

	m := map[string]struct{}{
		"foo": {},
		"bar": {},
	}
	fmt.Println("Keys only:")
	for x, _ := range m {
		fmt.Println(x)
	}
}

// ENDMAIN OMIT
