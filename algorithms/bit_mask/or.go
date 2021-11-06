package bitmask

import (
	"fmt"
)

/*
		Given operands a, b
		OR(a, b) = 1; when a = 1 or b = 1
	              else = 0

		We can use the nature of the bitwise OR operator
		to selectively set individual bits for a given
		integer. For instance, in the following example
		we use the OR operator to set (from least to most
		significant bits (MSB)) the 3rd, 7th, and 8th bit to 1.

*/

func OrOperator() {

	var a uint8 = 0
	a |= 196
	fmt.Printf("%b\n", a) //  11000100

	// useful for bitmasking techniques :
	a |= 3
	fmt.Printf("%b", a) //  11000111

}
