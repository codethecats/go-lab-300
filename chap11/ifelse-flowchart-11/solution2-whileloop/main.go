package main

import "fmt"

func main() {
	sumx := 0
	n := 0
	i := 1

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)

	for i <= n {
		sumx += i
		i++
	}
	fmt.Println(sumx)
}
