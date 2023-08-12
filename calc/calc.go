package calc

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func Minus(a int, b int) int {
	return a - b
}

func calcMsg(s string) string {
	fmt.Println("hello fmt")
	return "hi_1 " + s
}
