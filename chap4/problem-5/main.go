package main

import (
	"fmt"
	"reflect"
)

func main() {
	var numeric1 float64 = 5.0
	x := reflect.TypeOf(numeric1)

	fmt.Printf("%.1f %s \n", numeric1, x)
}
