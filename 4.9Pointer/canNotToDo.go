package main

import "fmt"

const i = 5

//func main()  {
//	ptr := &i
//	ptr2 := &10
//}
var p *int = nil

func main() {
	*p = 0
	fmt.Print(*p)
}
