package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var inputReader *bufio.Reader
	var sentence string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please insert sentence: ")
	sentence, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", sentence)
	}

	sentenceSlice := strings.Split(sentence, " ")

	fmt.Printf("%q\n", sentenceSlice)
}
