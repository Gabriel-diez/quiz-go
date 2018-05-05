package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"github.com/davecgh/go-spew/spew"
)

type quiz struct {
	question, solution string
}

func main() {
	l, err := readCSV()

	if err != nil {
		log.Fatal(err)
	}

	q := buildQuiz(l)

	spew.Dump(q)
}

func readCSV() ([][]string, error) {
	csvFile := flag.String("csv", "problems.csv", "Csv file in 'question, answer' format")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		return nil, fmt.Errorf("Couldn't open file")
	}

	r := csv.NewReader(file)

	l, err := r.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("Couldn't read file")
	}

	return l, nil
}

func buildQuiz(lines [][]string) []quiz {
	q := make([]quiz, len(lines))

	for i, l := range lines {
		q[i] = quiz {
			question: l[0],
			solution: l[1],
		}
	}
	return q
}