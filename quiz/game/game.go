package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/internal/question"
	"strings"
)

func Run(questions []question.Question) (correctAnswers uint) {
	fmt.Println("Welcome to the Country Quiz!")

	for _, q := range questions {
		if askQuestion(q) {
			correctAnswers++
		}
	}

	return correctAnswers
}

func askQuestion(q question.Question) bool {
	fmt.Printf("\nEnter the capital of %s: ", q.Country)

	if getUserInput() == strings.ToLower(q.Capital) {
		fmt.Println("\nCorrect!")
		return true
	} else {
		fmt.Printf("\nIncorrect capital of %s: ", q.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("\nAn error occurred while reading input. Try again.")
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\r\n"))
	}
}
