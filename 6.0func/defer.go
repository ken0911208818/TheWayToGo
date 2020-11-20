package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	a()
	f()
	test1("go")
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

func test1(s string) (n int, err error) {
	defer func() {
		log.Printf("test1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
