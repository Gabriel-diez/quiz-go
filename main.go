package main

import (
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"log"
	"strings"
	"time"
)

type quiz struct {
	question, solution string
}

func main() {
	csvFile := flag.String("csv", "problems.csv", "Csv file in 'question, answer' format")
	sec := flag.Int("time", 30, "Time in second for answer the quiz")
	flag.Parse()

	l, err := readCSV(*csvFile)

	if err != nil {
		log.Fatal(err)
	}

	quiz := buildQuiz(l)

	fmt.Printf("Press enter to start the quiz\n")
	fmt.Scanln()

	score := 0
	timer := time.NewTimer(time.Duration(*sec) * time.Second)

	quizLoop:
		for i, q := range quiz {
			fmt.Printf("Question %d: %s ?\n", i + 1, q.question)

			answerC := make(chan string)

			go func() {
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerC <- answer
			}()

			select {
			case <- timer.C:
				fmt.Printf("Time out !\n")
				break quizLoop;
			case answer := <-answerC:
				if answer == q.solution {
					score++
				}
			}
		}
	fmt.Printf("Your score is %d/%d.\n", score, len(quiz))
}

func readCSV(csvFile string) (lines [][]string, err error) {
	file, err := os.Open(csvFile)
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