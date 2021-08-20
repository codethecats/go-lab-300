package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	string1 := ""

	fmt.Print("Enter a Character: ")
	fmt.Scanf("%s", &string1)

	word := []rune(string1)

	fmt.Printf("%s %d %d \n", string1, utf8.RuneCountInString(string1), len(word))
}
