package bitmask

import "fmt"

/*

	Given operands a, b
	XOR(a, b) = 1; only if a != b
    else = 0

	Two integers a, b have the same signs when (a ^ b) ≥ 0 (or (a ^ b) < 0

*/

func XorOperator() {
	a, b := -12, 25
	fmt.Println("a and b have same sign?", (a^b) >= 0)
}
