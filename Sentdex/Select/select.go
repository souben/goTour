package main

import (
	"fmt"
)

func g1(c chan string) {
	c <- "Hello from chan 1"
}

func g2(c chan string) {
	c <- "Hello from chan 2"
}

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go g1(chan1)
	go g2(chan2)

	select {
	case v1 := <-chan1:
		fmt.Println(v1)
	case v2 := <-chan2:
		fmt.Println(v2)
	}

}
