package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"tutorial/greetings"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	log.SetPrefix("greeting: ")
	log.SetFlags(0)
	greetMsg, err := greetings.Greet("sasha")
	if err != nil {
		log.Fatal("BAD NAME")
	}
	fmt.Println(greetMsg)

	names := []string{"sasha", "anna", "roma"}

	for _, oneGreeting := range greetings.GreetMany(names) {
		fmt.Println(oneGreeting)
	}
}
