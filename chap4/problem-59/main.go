package main

import (
	"fmt"
	"sort"
)

func main() {

	ints := []int{18, 21, 19}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	s := []int{}
	s = sort.Ints(ints)
	fmt.Println("Sorted: ", s)
}
