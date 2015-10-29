package main

import (
	"fmt"
	"runtime"
	"sync"
)

func actor(id int, moo chan struct{}, meow chan struct{}, quit <-chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-moo:
			fmt.Println(id, "Moo!")
		case <-meow:
			fmt.Println(id, "Meow!")
		case <-quit:
			wg.Done()
			return
		}
	}
}

// ENDACTOR OMIT

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	moo := make(chan struct{})
	meow := make(chan struct{})
	quit := make(chan struct{})

	workers := 3
	var wg sync.WaitGroup
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go actor(i, moo, meow, quit, &wg)
	}

	commands := []chan struct{}{moo, meow, moo, moo, meow, moo, meow, meow}
	for _, c := range commands {
		c <- struct{}{}
	}

	close(quit)
	wg.Wait()
}

// ENDMAIN OMIT
