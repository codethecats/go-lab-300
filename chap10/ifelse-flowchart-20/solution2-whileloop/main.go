package main

import "fmt"

func main() {
	i := 14
	for i >= 8 {
		j := 5
		for j >= -3 {
			fmt.Println(i,j)
			j--
		}
		i--
	}
}
