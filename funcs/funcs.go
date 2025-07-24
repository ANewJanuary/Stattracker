package funcs

import "fmt"

func Greet(name string) string {
	var message string
	message = fmt.Sprintf("%v yo", name)
	return message
}

