package main

import "fmt"

var intP *int

func main() {
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
