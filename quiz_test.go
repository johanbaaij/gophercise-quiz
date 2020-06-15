package main

import (
	"reflect"
	"testing"
)

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

func TestTally(t *testing.T) {
	quiz := Quiz{[]Problem{{"1+1", "2", "2"}, {"2+2", "4", "5"}}}

	got := quiz.Tally()
	want := Tally{Correct: 1, Incorrect: 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
