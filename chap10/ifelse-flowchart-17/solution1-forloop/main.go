package main

import "fmt"

func main() {
	for i := -5; i <= 4; i += 2 {
		for j := 2; j <= 11; j += 3 {
			fmt.Println(i,j)
		}
	}
}
