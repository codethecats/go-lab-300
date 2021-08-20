package main

import (
	"fmt"
	"reflect"
)

func main() {
	var string1 string = "Golang"
	a := reflect.TypeOf(string1)
	fmt.Printf("%s %s \n", string1, a)
}
