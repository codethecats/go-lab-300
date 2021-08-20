package main

import "fmt"

func main() {
	list2 := []interface{}{"ant", "bird", "cat", "dog", "eagle"}

	list2[3] = "duck"

	fmt.Printf("%#v\n", list2)
}
