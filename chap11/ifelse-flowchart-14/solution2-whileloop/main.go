package main

import (
	"fmt"
	"math"
)

func main() {
	var sumx float64
	var n float64
	var i float64 = 1

	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &n)

	//Pow() function will return x**y, the base-x exponential of y.
	//var exponent, base float64
	//output := math.Pow(base, exponent)

	for i <= n {
		sumx += 1/math.Pow(i, 2)
		i++
	}

	fmt.Printf("%.02f \n", sumx)
}
