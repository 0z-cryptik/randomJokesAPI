package functions

import (
	"math/rand"
	"github.com/0z-cryptik/randomJokesAPI/jokes"
)

func GetRandomJoke() string {
	randomIndex := rand.Intn(len(jokes.Jokes))
	return jokes.Jokes[randomIndex]
}