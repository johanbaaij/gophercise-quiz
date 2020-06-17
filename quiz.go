package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
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

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	timeout := flag.Duration("timeout", 30*time.Second, "timeout in time.Duration format")
	flag.Parse()
	sleeper := &ConfigurableSleeper{*timeout, time.Sleep}

	file, _ := os.Open("problems.csv")
	problems := LoadProblems(file)
	quiz := Quiz{problems}

	fmt.Println("Press Enter to start the quiz")
	fmt.Scanln()
	fmt.Printf("Good luck! You've got %v\n", timeout)

	go QuizTimeLimit(os.Stdout, sleeper, quiz)

	for i, problem := range quiz.Problems {
		problem.PrintQuestion(os.Stdout)
		problem := problem.GetUserAnswer()
		quiz.Problems[i] = problem
	}

	quiz.PrintResults(os.Stdout)
}

// QuizTimeLimit exits the program after specified amount of time
func QuizTimeLimit(out io.Writer, sleeper Sleeper, quiz Quiz) {
	sleeper.Sleep()
	quiz.PrintResults(out)
	fmt.Fprintln(out, "Oops! Time's up.")
	os.Exit(0)
}

// LoadProblems converts a CSV into a slice of Problem structs
func LoadProblems(problemsCsv io.Reader) []Problem {
	reader := csv.NewReader(problemsCsv)
	lines, _ := reader.ReadAll()
	problems := make([]Problem, len(lines))

	for i, line := range lines {
		problems[i] = Problem{line[0], line[1], ""}
	}

	return problems
}

// PrintResults outputs the tally in readable format
func (q *Quiz) PrintResults(out io.Writer) {
	tally := q.Tally()
	fmt.Fprintf(out, "correct: %d\n", tally.Correct)
	fmt.Fprintf(out, "incorrect: %d\n", tally.Incorrect)
}

// PrintQuestion outputs the question
func (p Problem) PrintQuestion(out io.Writer) {
	fmt.Fprint(out, p.Question+"=\n")
}

// GetUserAnswer collects the answer from Stdin
func (p Problem) GetUserAnswer() Problem {
	var answer string
	fmt.Scanf("%s\n", &answer)
	p.RecordAnswer(answer)
	return p
}

// RecordAnswer stores a string in Problem.UserAnswer
func (p *Problem) RecordAnswer(answer string) {
	p.UserAnswer = answer
}

// Tally calculates the number of correct/incorrect answers
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

// Correct checks if an answer is correct
func (p Problem) Correct() bool {
	return (p.Answer == p.UserAnswer)
}
