package haha

import (
	"fmt"
	"github.com/ahd99/webTest/calc"

	"rsc.io/quote"
)

func SayHaHa() {
	fmt.Printf("Hello golang\n")
	fmt.Printf("%s\n", quote.Hello())
	fmt.Printf("%d\n", calc.Sum(5, 12))
}
