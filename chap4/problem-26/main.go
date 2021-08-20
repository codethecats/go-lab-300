package main

import (
	"fmt"
)

func main() {
	var string1 = "Golang"
	fmt.Printf("%c \n", []rune(string1)[1:4])
}
