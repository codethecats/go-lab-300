package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Interger Reverse Sort")
	list4 := []int{0, 4, 2, 3, 1}
	sort.Sort(sort.Reverse(sort.IntSlice(list4)))
	fmt.Println(list4)
}
