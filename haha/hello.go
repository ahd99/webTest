package main

import (
	"fmt"
	"github.com/ahd99/webTest/haha/calc"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("Hello golang\n")
	fmt.Printf("%s\n", quote.Hello())
	fmt.Printf("%d\n", calc.Sum(5, 7))
}