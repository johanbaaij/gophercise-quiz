package main

import "testing"

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
