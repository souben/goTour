package bitmask

import "fmt"

/*

	Given integer operands a and n,
	a << n; shifts all bits in a to the left n times
	a >> n; shifts all bits in a to the right n times

*/

func LeftRightShift() {
	var a int8 = 3
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a<<1)
	fmt.Printf("%08b\n", a<<2)
	fmt.Printf("%08b\n", a<<3)

	// divide by 2*b
	fmt.Println(200 >> 1)

	// multiple by 2*b
	fmt.Println(12 << 2)

	// the | and << operators are used to set the n rd bit

	var b int8 = 8
	b = b | (1 << 2)
	fmt.Printf("%08b\n", b)

	// the & operators to test if nth bit is set in a value
	if a&(1<<2) != 0 {
		fmt.Println("take action")
	}

	// the &^ and the shift operators, we can unset the nth bit of a value
	a = a &^ (1 << 2)
	fmt.Printf("%04b\n", a)

}
