package main

import (
	"fmt"
)

func main() {
	string1 := "Golang"
	convertNegativeNumberIntoPositive1 := -4
	convertNegativeNumberIntoPositive2 := -2

	word := []rune(string1)
	result := string1[convertNegativeNumberIntoPositive1+len(word) : convertNegativeNumberIntoPositive2+len(word)+1]

	fmt.Printf("%v\n", result)

}
