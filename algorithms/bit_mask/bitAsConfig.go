package bitmask

import (
	"fmt"
	"strings"
)

/*

	Now, recall that AND(a, 1) = a if and only if a = 1. We can use that fact to query a value for its set bits.
	For instance, from the code above a & 196 will return 196 because the bits for that value are indeed set in a.
	So we can combine the use of the OR and the AND as a way of specifying configuration values and reading them respectively.
	The following source code snippet shows this at work. Function procstr transforms the content of a string.
	It takes two parameters: the first parameter, str, is the string to be transformed and the second parameter,
	conf, is an integer used to specify multiple transformation configurations using bit masking.

*/

// IOTA : https://golangbyexample.com/iota-in-golang/

const (
	UPPER = 1 << iota
	LOWER
	CAP
	REV
)

func CombineAndOr() {
	str := procstr("HOla", REV|LOWER|CAP)
	fmt.Println(str)
}

func reverse(str string) string {
	runes := []rune(str)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func procstr(str string, conf byte) string {

	if conf&UPPER != 0 {
		str = strings.ToUpper(str)
	}

	if conf&REV != 0 {
		str = reverse(str)
	}

	if conf&LOWER != 0 {
		str = strings.ToLower(str)
	}
	if conf&CAP != 0 {
		str = strings.Title(str)
	}
	return str
}
