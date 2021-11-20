package partone

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var flagvar int

func init() {
	flag.IntVar(&flagvar, "n", 2, "number of questions to answer")
}

func PartOne() {
	flag.Parse()
	allrecords := csvParser("problems.csv")

	count := 0
	i := 0
	for i < flagvar {
		fmt.Printf("what is %s ?", allrecords[i][0])
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		if strings.Trim(str, "\n") == allrecords[i][1] {
			count++
		}
		i++
	}
	fmt.Printf("you scored %d of %d questions\n", count, flagvar)
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
