package main

import "fmt"

type A struct {
	name string
}

func main() {
	a := A{
		name: "ken",
	}
	fmt.Printf("%p\n", &a)
	DoSomething(&a)
	DoSomething1(a)
}

func DoSomething(a *A) {
	fmt.Printf("%p\n", a)
	b := a
	fmt.Print(b, &b)
}

func DoSomething1(a A) {
	b := &a
	fmt.Println(b, &b)
}
