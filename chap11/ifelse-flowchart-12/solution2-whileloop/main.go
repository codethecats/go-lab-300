package main

import "fmt"

func main() {
	productx := 1
	n := 0
	i := 1

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)

	for i <= n {
		productx *= i
		i++
	}

	fmt.Println(productx)
}
