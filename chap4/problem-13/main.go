package main

import (
	"fmt"
	"reflect"
)

func main() {
	var logic2 bool = false
	x := reflect.TypeOf(logic2)
	fmt.Printf("%t %s \n", logic2, x)
}
