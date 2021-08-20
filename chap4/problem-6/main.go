package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numeric2 float64 = -3.1
	x := reflect.TypeOf(numeric2)

	fmt.Printf("%.1f %s \n", numeric2, x)
}
