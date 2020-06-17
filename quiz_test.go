package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestLoadProblems(t *testing.T) {
	csv := `1+1,2
2+2,4`

	want := []Problem{{"1+1", "2", ""}, {"2+2", "4", ""}}
	got := LoadProblems(strings.NewReader(csv))

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestPrintQuestion(t *testing.T) {
	problem := Problem{"1+1", "2", ""}
	buffer := &bytes.Buffer{}
	problem.PrintQuestion(buffer)
	got := buffer.String()
	want := "1+1=\n"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestRecordAnswer(t *testing.T) {
	problem := Problem{"1+1", "2", ""}
	answer := "2"
	problem.RecordAnswer(answer)
	want := problem.UserAnswer

	if answer != want {
		t.Errorf("got %v want %v", answer, want)
	}
}

func TestAnswerCorrectness(t *testing.T) {
	problem := Problem{"1+1", "2", "2"}

	if problem.Correct() != true {
		t.Errorf("got %v want %v", problem.Correct(), true)
	}
}

func TestShowResults(t *testing.T) {
	quiz := Quiz{[]Problem{{"1+1", "2", "2"}, {"2+2", "4", "5"}}}

	got := quiz.Tally()
	want := Tally{Correct: 1, Incorrect: 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
