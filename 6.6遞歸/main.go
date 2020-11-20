package main

import "fmt"

func main() {
	result := 0
	for i := 0; i <= 1000000; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	//fmt.Println("-------------------")
	//for j := 10; j>0; j-- {
	//	result = fibonacci(j)
	//	fmt.Printf("fibonacci(%d) is: %d\n", j, result)
	//}
	//fmt.Println("-------------------")
	//for i := 0; i <= 30; i++ {
	//	result = homeWork(i)
	//	fmt.Printf("homeWork(%d) is: %d\n", i, result)
	//}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

func homeWork(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = n * homeWork(n-1)
	}
	return
}
