package main

import (
	"fmt"
	"reflect"
)

func main() {
	var logic1 bool = true
	x := reflect.TypeOf(logic1)

	fmt.Printf("%t %s \n", logic1, x)
}
