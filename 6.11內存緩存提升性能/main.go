package main

import (
	"fmt"
	"sync"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		var result uint64 = 0
		start := time.Now()
		for i := LIM - 1; i >= 0; i-- {
			result = testFibonacci(i)
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		}
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	}()
	go func() {
		defer wg.Done()
		var result uint64 = 0
		start := time.Now()
		for i := LIM - 1; i >= 0; i-- {
			result = fibonacci(i)
			fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		}
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("longCalculation took this amount of time: %s\n", delta)
	}()
	wg.Wait()
	fmt.Println("finish")
}

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}

func testFibonacci(n int) (res uint64) {
	if n <= 1 {
		res = 1
		return
	} else {
		res = testFibonacci(n-1) + testFibonacci(n-2)
		return
	}
}
