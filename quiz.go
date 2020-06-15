package main

import (
	"fmt"
	"io"
)

type Quiz struct {
	Problems []Problem
}

type Tally struct {
	Correct   int
	Incorrect int
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

type Problem struct {
	Question   string // a question
	Answer     string // the correct answer
	UserAnswer string // answer provided by user
}

func (p *Problem) RecordAnswer(answer string) {
	p.UserAnswer = answer
}

func (p Problem) DisplayQuestion(out io.Writer) {
	fmt.Fprint(out, p.Question+"=")
}

func (p Problem) Correct() bool {
	return p.Answer == p.UserAnswer
}

func main() {
	// load csv
	// form quiz
	// start loop quiz
	// display question
	// collect answer
	// ask questions (done)
	// display tally (done)
}
