package main

import (
	"fmt"
	"strings"
)

func main() {
	strings1 := ""

	fmt.Print("Please insert string: ")
	fmt.Scanf("%s", &strings1)

	if strings.Contains(strings1, "ก") == true {
		fmt.Printf("%t \n", true)
	} else {
		fmt.Printf("%t \n", false)
	}
}
่