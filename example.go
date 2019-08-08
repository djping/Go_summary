package main

import (
	"fmt"
)
func main () {
	fmt.Println("this is my first Golang demo!")
	var a int = 5
	b := 20
	if a > b {
		fmt.Println("a>b")
	} else {
		fmt.Println("a<b")
	}
	for i:=0; i < a; i++ {
		fmt.Println("add_func: ", add(1, 2))
	}
}
func add (n1, n2 int) int {
	return n1 + n2
}