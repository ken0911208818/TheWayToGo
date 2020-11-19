package main

import "fmt"

func main() {
	s := "hello word"
	var p *string
	p = &s
	*p = "你好"
	fmt.Printf("原本的s :%s 記憶體位置 : %p 被更改過的Ｐ: %s, 記憶體位置 ： %p", s, &s, *p, p)
}
