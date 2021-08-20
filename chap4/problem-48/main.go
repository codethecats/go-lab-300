package main

import "fmt"

func main() {
	list1 := []interface{}{0, 1, 2, "a", "b", "c"}

	fmt.Printf("%#v\n", list1[3:6])
}
