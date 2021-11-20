package parttwo

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

var (
	flagvar               int
	flagTimeLimitQuiz     int
	flagTimeLimitQuestion int
)

func init() {
	flag.IntVar(&flagvar, "n", 7, "number of questions to answer")
	flag.IntVar(&flagTimeLimitQuiz, "t", 20, "time limit to answer all questions")
	flag.IntVar(&flagTimeLimitQuestion, "T", 2, "time limit to answer a question")
}

func PartTwo() {
	flag.Parse()
	allrecords := csvParser("problems.csv")
	duration := time.NewTicker(time.Duration(flagTimeLimitQuestion) * time.Second)
	timeLimit := time.NewTimer(time.Duration(flagTimeLimitQuiz) * time.Second)
	count := 0
	i := -1
	var replyChan = make(chan string, flagvar)
	for {

		select {
		case <-duration.C:
			i++
			fmt.Printf("what is %s ?\n", allrecords[i][0])
			go reply(replyChan)

		case <-timeLimit.C:
			fmt.Printf("\nyou scored %d of %d questions\n", count, flagvar)
			return
		case str := <-replyChan:
			fmt.Println(allrecords[i][1], strings.Trim(string(str), "\n") == allrecords[i][1], string(str))
			if strings.Trim(string(str), "\n") == allrecords[i][1] {
				count++
			}
			if i == flagvar-1 {
				fmt.Printf("you scored %d of %d questions\n", count, flagvar)
				return
			}
		}
	}

}

func reply(ch chan string) {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	ch <- str
}

func csvParser(f string) [][]string {

	var records [][]string
	s, err := os.Open(f)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(s)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, record)
	}
	return records
}
