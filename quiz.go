package main

type Problem struct {
	Question   string // a question
	Answer     string // the correct answer
	UserAnswer string // answer provided by user
}

func (p *Problem) RecordAnswer(answer string) {
	p.UserAnswer = answer
}

func (p Problem) Correct() bool {
	return p.Answer == p.UserAnswer
}

func main() {

}
