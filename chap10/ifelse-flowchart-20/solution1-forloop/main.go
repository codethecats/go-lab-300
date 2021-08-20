package main

import "fmt"

func main() {
	for i := 14; i >= 8; i-- {
		for j := 5; j >= -3; j-- {
			fmt.Println(i,j)
		}
	}
}
