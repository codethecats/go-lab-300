package main

import "fmt"

func main() {
	list1 := []interface{}{0, 1, 2, "a", "b", "c"}

	convertNegativeNumberIntoPositive := -3

	result := list1[convertNegativeNumberIntoPositive+len(list1)]

	fmt.Printf("%v\n", result)
}
