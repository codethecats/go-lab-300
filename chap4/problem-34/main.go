package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := ""
	string2 := ""

	fmt.Print("string1: ")
	fmt.Scanf("%s", &string1)

	fmt.Print("string2: ")
	fmt.Scanf("%s", &string2)

	if strings.Contains(string1, string2) == true {
		fmt.Printf(" %t \n", true)
	} else {
		fmt.Printf(" %t \n", false)
	}
}
