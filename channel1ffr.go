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
	go func(c chan int) {
		first := <-c
		fmt.Println("First:", first)

		second := <-c
		fmt.Println("Second:", second)
		fmt.Println("Meaning of life, universe and everything:", first*second)

		wg.Done()
	}(nums)

	nums <- 6
	nums <- 7
	wg.Wait()
}

// ENDMAIN OMIT
