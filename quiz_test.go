package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
	"time"
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

func TestResults(t *testing.T) {
	quiz := Quiz{[]Problem{{"1+1", "2", "2"}, {"2+2", "4", "5"}}}

	t.Run("test Tally() method", func(t *testing.T) {
		got := quiz.Tally()
		want := Tally{Correct: 1, Incorrect: 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("test print output", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		quiz.PrintResults(buffer)
		got := buffer.String()
		want := `correct: 1
incorrect: 1
`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
