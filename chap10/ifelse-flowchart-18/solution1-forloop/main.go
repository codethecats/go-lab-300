package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		for j := 4; j >= -6; j-- {
			fmt.Println(i,j)
		}
	}
}
