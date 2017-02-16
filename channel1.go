// +build OMIT

package main

import "fmt"

func main() {
	nums := make(chan int)

	go func() {
		first := <-nums
		fmt.Println("Multiplicand:", first)

		second := <-nums
		fmt.Println("Multiplier:", second)
		fmt.Println("Meaning of life, universe and everything: ", first*second)
	}()

	nums <- 6
	nums <- 7
}

// ENDMAIN OMIT
