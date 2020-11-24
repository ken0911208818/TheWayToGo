package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]
	// load the array with integers: 0,1,2,3,4,5
	for index := range arr1 {
		arr1[index] = index
	}
	// print the slice

	for index, value := range slice1 {
		fmt.Println("Slice item", index, "is :", value)
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//grow the slice

	slice1 = slice1[:4]
	for index := range slice1 {
		fmt.Println(`Slice at `, index, `is`, slice1[index])
	}
	Q1()
}

func Q1() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Printf("b[1:4] is (%s)\n", b[1:4])
	fmt.Printf("b[:4] is (%s)\n", b[:2])
	fmt.Printf("b[2:] is (%s)\n", b[2:])
	fmt.Printf("b[:] is (%s)\n", b[:])
}
