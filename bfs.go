// +build OMIT

package main

import "fmt"
import "math"
import "gopkg.in/eapache/queue.v1"

func bfs(graph map[uint32][]uint32, start uint32, size uint32, fn func(id uint32) bool) (node uint32, found bool) {
	nodes := make(chan uint32, size)

	type Ret struct {
		node  uint32
		found bool
	}
	ret := make(chan Ret)

	for {
		select {
		case node := <-nodes:
			if fn(node) {
				ret <- Ret{node, true}
			}
			for _, child := range graph[node] {
				nodes <- child
			}
		default:
			ret <- Ret{math.MaxUint32, false}
		}
	}

	return <-ret
}

func main() {
	q := queue.New()
	q.Add(1)
	fmt.Println(q.Peek())
}
