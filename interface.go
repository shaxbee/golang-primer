// +build OMIT

package main

import "fmt"

type Animal interface {
	Talk()
}

type Cat struct {
	Name string
}

func (c Cat) Talk() {
	fmt.Println(c.Name, "says Meow")
}

func main() {
	var animal Animal = Cat{"Charlie"}
	animal.Talk()
}
