package main

import (
	"bytes"
	"fmt"
)

type Buffer struct {
	name string
	age  int
}

func main() {
	var buffer bytes.Buffer

	buffer.WriteString("hello World")

	fmt.Println(buffer.String(), "\n")
}
