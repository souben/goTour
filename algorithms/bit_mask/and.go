package bitmask

import (
	"fmt"
	"math/rand"
)

/*
	Given operands a, b
	AND(a, b) = 1; only if a = b = 1
					else = 0

*/

func AndOperator() {

	for x := 0; x < 5; x++ {
		num := rand.Int()
		fmt.Printf("%d,  %d\n", num, num&1)
	}

}
