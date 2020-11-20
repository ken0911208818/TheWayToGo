package main

import "fmt"

func main() {
	Greetings("Name: ", "joe", 123, "ken")
}

func Greeting(prefix string, arg ...string) {
	if len(arg) == 0 {
		fmt.Println("is error for no arg")
		return
	}
	for _, val := range arg {
		fmt.Println(prefix, val)
	}
	return
}

func Greetings(prefix string, values ...interface{}) {
	for _, val := range values {
		switch val.(type) {
		case int:
			fmt.Printf(prefix, "is int%d\n", val)
		case string:
			fmt.Printf(prefix, "is string%s\n", val)
		default:
			fmt.Printf(prefix, "not any")
		}

	}
}
