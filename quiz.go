package main

import (
	"encoding/csv"
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

func main() {
	file, _ := os.Open("problems.csv")
	problems := LoadProblems(file)
	quiz := Quiz{problems}

	for i, problem := range quiz.Problems {
		problem.PrintQuestion(os.Stdout)
		problem := problem.GetUserAnswer()
		quiz.Problems[i] = problem
	}

	quiz.ShowResults()
}

// Converts a CSV into a slice of Problem structs
func LoadProblems(problemsCsv io.Reader) []Problem {
	reader := csv.NewReader(problemsCsv)
	lines, _ := reader.ReadAll()
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{line[0], line[1], ""}
	}

	return problems
}

func (q *Quiz) ShowResults() {
	tally := q.Tally()
	fmt.Printf("correct: %d\n", tally.Correct)
	fmt.Printf("incorrect: %d\n", tally.Incorrect)
}

func (p Problem) PrintQuestion(out io.Writer) {
	fmt.Fprint(out, p.Question+"=\n")
}

func (p Problem) GetUserAnswer() Problem {
	var answer string
	fmt.Scanf("%s\n", &answer)
	p.RecordAnswer(answer)
	return p
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
	return (p.Answer == p.UserAnswer)
}
