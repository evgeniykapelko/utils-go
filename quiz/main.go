package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/internal/question"
	"quiz/shuffler"
)

func main() {
	questions, err := question.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load questions: %v\n", err)
		os.Exit(1)
	}

	shuffler.Shuffle(questions)

	correctAnswers := game.Run(questions)

	fmt.Printf("You scored %d out of %d!\n", correctAnswers, len(questions))
}
