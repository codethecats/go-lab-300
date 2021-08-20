package main

import (
	"fmt"
	"strings"
)

func main() {
	var string1 string
	fmt.Println("Enter 1st String")
	fmt.Scanf("%s", &string1)

	string1Slice := strings.Split(string1, "")

	fmt.Println(string1Slice)
}
