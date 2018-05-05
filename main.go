package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

type quiz struct {
	question, solution string
}

func main() {
	l, err := readCSV()

	if err != nil {
		log.Fatal(err)
	}

	quiz := buildQuiz(l)

	score := 0

	for i, q := range quiz {
		fmt.Printf("Question %d: %s ?\n", i + 1, q.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == q.solution {
			score++
		}
	}

	fmt.Printf("Your score is %d/%d.\n", score, len(quiz))

}

func readCSV() (lines [][]string, err error) {
	csvFile := flag.String("csv", "problems.csv", "Csv file in 'question, answer' format")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		return nil, fmt.Errorf("Couldn't open file")
	}

	defer file.Close()

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
			solution: strings.TrimSuffix(l[1], ""),
		}
	}
	return q
}