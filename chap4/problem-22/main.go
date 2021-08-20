package main

import (
	"fmt"
)

func main() {
	string1 := "Golang"
	convertNegativeNumberIntoPositive := -6

	word := []rune(string1)
	result := string1[convertNegativeNumberIntoPositive + len(word)]

	fmt.Printf("%#c\n", result)
}
