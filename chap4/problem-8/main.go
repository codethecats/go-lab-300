package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	x := reflect.TypeOf(input)

	fmt.Printf("%.1f %s \n", input, x)
}
