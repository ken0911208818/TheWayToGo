package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	start := time.Now()
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	where()
	// make a special Adder function, a gets value 2:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))
	var f = f()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
	}(1000) // Passes argument 1000 to the function literal.
	time.Sleep(2 * time.Second)
	fmt.Println(g)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("main is of time: %s\n", delta)
}

func Add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

func f() func(int) int {
	var i int
	return func(delta int) int {
		i += delta
		return i
	}
}
