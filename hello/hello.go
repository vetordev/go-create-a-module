package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	name := []string{"V", "JC", "Vi"}
	messages, err := greetings.Hellos(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
