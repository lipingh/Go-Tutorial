package greetings

import "fmt"

func Hello(name string) string {
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome", name)
	// the following is shortcut for declaring and initializing a variable
	message := fmt.Sprintf("Hi, %v. Welcome", name)
	return message
}
