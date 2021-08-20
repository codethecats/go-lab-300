package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanf("%f", &input)
	x := reflect.TypeOf(input)
	fmt.Printf("%d %s \n", input, x)
}
