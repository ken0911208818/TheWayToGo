package main

import "fmt"

func main() {
	var arr1 [5]int
	for k, _ := range arr1 {
		arr1[k] = k * 2
	}
	Q1()
	arr2 := &arr1
	arr2[2] = 100
	for _, val := range arr1 {
		fmt.Printf(`the number is %d\n`, val)
	}
}

func Q1() {
	a := [...]string{"a", "b", "c", "d", "e"}
	for i := range a {
		fmt.Println(`Array item`, i, "is", a[i])
	}
}
