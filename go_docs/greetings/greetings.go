package greetings

import (
	"errors"
	"log"
	"math/rand"
)

func GreetMany(names []string) []string {
	var result []string
	for _, name := range names {
		res, err := Greet(name)
		if err != nil {
			log.Fatal("bad name")
		}
		result = append(result, res)
	}
	return result
}
func Greet(name string) (string, error) {

	if name == "" {
		return "", errors.New("")
	}
	message := randomGreeting() + name
	return message, nil
}

func randomGreeting() string {

	formats := [3]string{"hello ", "welcome ", "hi "}

	return formats[rand.Intn(len(formats))]
}
