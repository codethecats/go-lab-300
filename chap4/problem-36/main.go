package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := ""

	fmt.Print("Enter a Character: ")
	fmt.Scanf("%s", &string1)

	string2 := strings.Replace(string1, "a", "A", -1)

	fmt.Printf("%s \n", string2)
}
