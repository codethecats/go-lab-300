package main

import "fmt"

func main() {
	var sumx float64
	var n float64
	var i float64 = 1

	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &n)

	for i <= n {
		sumx += 1/i
		i++
	}

	fmt.Printf("%.02f \n", sumx)
}
