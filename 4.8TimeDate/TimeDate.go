package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())                              // 現在時間
	fmt.Printf("This time is : %d\n", time.Now().Unix()) //轉成Unix
}
