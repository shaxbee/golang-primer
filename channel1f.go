// +build OMIT
package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		first := <-nums
		fmt.Println("Multiplicand:", first)

		second := <-nums
		fmt.Println("Multiplier:", second)
		fmt.Println("Meaning of life, universe and everything:", first*second)

		wg.Done()
	}()

	nums <- 6
	nums <- 7
	wg.Wait()
}

// ENDMAIN OMIT
