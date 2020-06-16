package main

import (
	"fmt"
	"io"
	"os"
)

type Quiz struct {
	Problems []Problem
}

type Problem struct {
	Question   string // a question
	Answer     string // the correct answer
	UserAnswer string // answer provided by user
}

type Tally struct {
	Correct   int
	Incorrect int
}

func (q *Quiz) Run() {
	for _, problem := range q.Problems {
		problem.PrintQuestion(os.Stdout)
	}
	fmt.Println("done")
}

func (p Problem) PrintQuestion(out io.Writer) {
	fmt.Fprint(out, p.Question+"=\n")
}

func (p *Problem) RecordAnswer(answer string) {
	p.UserAnswer = answer
}

func (q Quiz) Tally() Tally {
	tally := Tally{Correct: 0, Incorrect: 0}

	for _, problem := range q.Problems {
		if problem.Correct() {
			tally.Correct++
		} else {
			tally.Incorrect++
		}
	}

	return tally
}

func (p Problem) Correct() bool {
	return p.Answer == p.UserAnswer
}

func main() {

}
