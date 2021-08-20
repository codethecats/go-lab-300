package main

import "fmt"

func main() {
	i := 0
	for i <= 5 {
		j := 4
		for j >= -6 {
			fmt.Println(i,j)
			j--
		}
		i++
	}
}
