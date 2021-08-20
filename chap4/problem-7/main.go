package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numeric3 float64 = 0.0
	x := reflect.TypeOf(numeric3)

	fmt.Printf("%.1f %s \n", numeric3, x)
}
