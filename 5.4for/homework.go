package main

import "fmt"

func main() {
	HW1()
	HW4()
}

func HW1() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}

func HW2() {
	for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}
}

func HW3() {
	for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}
}

func HW4() {
	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}
}
