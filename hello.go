package main

import (
	"fmt"
	"hd99/webTest/calc"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello golang\n")
	fmt.Printf("%s\n", quote.Hello())
	fmt.Printf("%d\n", calc.Sum(5, 7))
}
