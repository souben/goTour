package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

var limitTime int

func init() {
	flag.IntVar(&limitTime, "limit", 2, "help message for flagname")
}

func main() {
	file, err := os.Open("problems.csv")
	reader := bufio.NewReader(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	q := make([][]string, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		q = append(q, record)
	}

	counter := 0
	for _, v := range q {
		c := make(chan string)
		flag.Parse()
		fmt.Printf("what %s, sir? :", v[0])

		go func() {
			res, _ := reader.ReadString('\n')
			c <- strings.TrimSpace(res)
		}()

		select {
		case m := <-c:
			if v[1] == m {
				counter++
			}
		case <-time.After(2 * time.Second):
			fmt.Println("timed out")
			break
		}
	}

	fmt.Printf("you scored %d of %d\n", counter, len(q))
}
