package main

import "fmt"

func main() {
	productx := 1
	n := 0

	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		productx *= i
	}

	fmt.Println(productx)
}
