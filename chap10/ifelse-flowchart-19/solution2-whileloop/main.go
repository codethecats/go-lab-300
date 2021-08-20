package main

import "fmt"

func main() {
	i := 11
	for i >= 0 {
		j := -7
		for j <= 0 {
			fmt.Println(i,j)
			j++
		}
		i -= 2
	}
}
