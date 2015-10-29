// +build OMIT

package main

import "fmt"

func foo(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func main() {
	for i := range foo(10) {
		fmt.Println(i)
	}
}
