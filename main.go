package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var splitString []string
	lowerString := strings.ToLower(text)
	lowerString = strings.TrimSpace(lowerString)
	splitString = strings.Fields(lowerString)
	fmt.Println(splitString)
	return splitString
}

func main() {
	testString := " hello world "
	cleanInput(testString)

	fmt.Println("This is here so I don't have to comment and uncomment fmt a billion times for debugging")
}
