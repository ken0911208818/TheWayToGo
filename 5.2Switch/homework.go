package main

import "fmt"

func main() {
	fmt.Printf(Season(3))
}

func Season(month int) string {
	switch month {
	case 12, 1, 2:
		return fmt.Sprintf("冬天")

	case 3, 4, 5:
		return fmt.Sprintf("春天")
	case 6, 7, 8:
		return fmt.Sprintf("夏天")
	case 9, 10, 11:
		return fmt.Sprintf("秋天")
	default:
		return fmt.Sprintf("沒有這個季節")
	}
}
