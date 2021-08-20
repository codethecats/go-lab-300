package main

import (
	"fmt"
)

func main() {
	string1 := "Golang"
	convertNegativeNumberIntoPositive := -2

	word := []rune(string1)
	result := string1[2 : convertNegativeNumberIntoPositive+len(word)+1]

	fmt.Printf("%v\n", result)
}