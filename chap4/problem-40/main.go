package main

import (
	"fmt"
	"reflect"
)

func main() {
	list1 := []interface{}{0, 1, 2, "a", "b", "c"}

	fmt.Println(list1)
	fmt.Println(reflect.TypeOf(list1))
}
