package main

import "fmt"

func main() {
	sumx := 0
	n := 0

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		sumx += i
	}

	fmt.Println(sumx)
}
