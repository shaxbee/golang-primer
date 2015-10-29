// +build OMIT

package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	a := []int{1, 2, 3, 4}
	for {
		a = a[1:]
		if len(a) == 0 {
			break
		}
		fmt.Println(a)
	}
}

// ENDMAIN OMIT
