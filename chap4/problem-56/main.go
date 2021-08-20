package main

import (
	"fmt"
	"sort"
)

func main() {
	list4 := []int{0, 4, 2, 3, 1}

	sort.Ints(list4)
	fmt.Println(list4)
}
