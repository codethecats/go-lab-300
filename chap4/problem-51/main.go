package main

import "fmt"

func main() {
	list2 := []interface{}{"ant", "bird", "cat", "dog", "eagle"}

	list2[0] = "ape"

	fmt.Printf("%#v\n", list2)
}
