package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := ""

	fmt.Print("Please insert string: ")
	fmt.Scanf("%s", &string1)

	fmt.Printf("%s \n", strings.Replace(string1, "a", "A", -1))
}
