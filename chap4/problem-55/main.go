package main

import "fmt"

func main() {
	list3 := []string{"apple", "banana", "cherry", "eggfruit"}
	convertNegativeNumberIntoPositive := -2

	list3 = append(list3, "new")
	copy(list3[3:], list3[2:])

	convertNegativeNumberIntoPositive = convertNegativeNumberIntoPositive + len(list3)

	fmt.Printf("%#v\n", convertNegativeNumberIntoPositive)
	list3[convertNegativeNumberIntoPositive] = "kiwi"

	fmt.Printf("%#v\n", list3)
}
