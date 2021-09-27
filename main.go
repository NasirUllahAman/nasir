package main

import (
	parent "family/parents"
	child "family/parents/child"

	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("Mr bawar khan"))
	s := new(child.Son)
	fmt.Println(s.Data("mr Aman Ullah"))
}
