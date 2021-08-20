package main

import "fmt"

func main() {
	list1 := []interface{}{0, 1, 2, "a", "b", "c"}

	convertNegativeNumberIntoPositive := -2

	result := list1[1:convertNegativeNumberIntoPositive+len(list1)+1]

	fmt.Printf("%v\n", result)
}
