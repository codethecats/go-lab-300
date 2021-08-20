package main

import (
	"bytes"
	"fmt"
)

func main() {
	sentence := ""
	c := ""

	fmt.Print("Please insert sentence: ")
	fmt.Scanf("%s", &sentence)

	fmt.Print("Please insert c: ")
	fmt.Scanf("%s", &c)

	splitExample := bytes.Split([]byte(sentence), []byte(c))

	joinChar := []byte{','}
	countryJoin := bytes.Join(splitExample, joinChar)

	fmt.Printf("Join joins slice with ',' : %q \n", countryJoin)
}
