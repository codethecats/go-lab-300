package main

import (
	"fmt"
)

func main() {
	list3 := []string{"apple", "cherry", "eggfruit"}

	list3 = append(list3, "new")
	copy(list3[2:], list3[1:])

	list3[1] = "banana"

	fmt.Println(list3)
}
