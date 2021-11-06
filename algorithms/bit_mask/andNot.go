package bitmask

import "fmt"

/*

Given operands a, b
AND_NOT(a, b) = AND(a, NOT(b))

*/

func AndNotOperator() {
	var a byte = 0xAB
	fmt.Printf("%08b\n", a)
	a &^= 0x0F
	fmt.Printf("%08b\n", a)

}
