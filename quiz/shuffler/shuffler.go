package shuffler

import (
	"math/rand"
	"quiz/internal/question"
)

func Shuffle(questions []question.Question) {
	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})
}
