package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	//if no name is given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// if name, was given, return the message
	message := fmt.Sprintf("Hi %v", name)
	return message, nil
}
