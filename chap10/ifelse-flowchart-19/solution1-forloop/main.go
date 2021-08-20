package main

import "fmt"

func main() {
	for i := 11; i >= 0; i -= 2 {
		for j := -7; j <= 0; j++ {
			fmt.Println(i,j)
		}
	}
}
