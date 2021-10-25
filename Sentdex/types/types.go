package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func mutiply(a, b string) (string, string) {
	return a, b
}

const x float32 = 5.5

func main() {
	num1, num2 := float32(0.9), float32(19.0)
	var result float32
	result = add(num1, num2)
	fmt.Println(result * x)

	w1, w2 := "Hello", "man"
	var phrase, phrase2 string
	phrase, phrase2 = mutiply(w1, w2)
	fmt.Println(phrase)
	fmt.Println(phrase2)
}
