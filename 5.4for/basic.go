package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	Practive1()
	Practive()
	character()
	homework1()
	homework2()
	homework3()
}

func Practive() {
	for i := 0; i < 15; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}

func Practive1() {
	var i = 0
start:
	fmt.Printf("This is the %d iteration\n", i)
	i++
	if i < 15 {
		goto start
	}
}

func character() {
	for i := 1; i <= 25; i++ {
		fmt.Println(strings.Repeat("G", i))
	}
}

func homework1() {
	for i := 0; i <= 10; i++ {
		// %b 取int 的二進位
		// ^i 取int 的-1補數
		fmt.Printf("the complement of %b is: %b\n", i, ^i)
	}
}

func homework2() {
	const (
		FIZZ     = 3
		BUZZ     = 5
		FIZZBUZZ = 15
	)
	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
func homework3() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			fmt.Print(strings.Repeat("*", 1))
		}
		fmt.Print("\n")
	}
}
