package main

import "fmt"

func main() {
	i := 10
	for i >= 1 {
		j := 1
		for j <= 5 {
			fmt.Println(i, j)
			j++
		}
		i -= 1
	}
}
