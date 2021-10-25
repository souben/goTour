package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	defer wg.Done()
}

func main() {
	wg.Add(1)
	go say("Hi")
	wg.Wait()
	wg.Add(1)
	go say("There")
	wg.Add(1)

	go say("Hello")
	wg.Wait()
}
