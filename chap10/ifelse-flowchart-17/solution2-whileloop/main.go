package main

import "fmt"

func main() {
	i := -5
	for i <= 4 {
		j := 2
		for j <= 11 {
			fmt.Println(i,j)
			j += 3
		}
		i += 2
	}
}
