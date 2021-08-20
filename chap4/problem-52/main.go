package main

import "fmt"

func main() {
	list2 := []interface{}{"ant", "bird", "cat", "dog", "eagle"}

	list2 = append(list2, "fish")

	fmt.Printf("%#v\n", list2)
}
