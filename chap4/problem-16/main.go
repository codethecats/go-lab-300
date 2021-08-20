package main

import (
	"fmt"
	"reflect"
)

func main() {
	complex1 := 1 + 2i
	fmt.Printf("Complex: %g Type of: %s \n", complex1, reflect.TypeOf(complex1))
}
