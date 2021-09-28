package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// // log out the date
	// log.SetFlags(1)
	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	// message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
