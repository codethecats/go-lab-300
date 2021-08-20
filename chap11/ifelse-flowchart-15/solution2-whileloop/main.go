package main

import (
	"fmt"
	"math"
)

func main() {
	var sumx float64
	var n float64
	var pi float64
	var i float64 = 1

	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &n)

	for i <= n {
		sumx += 1/math.Pow(i, 2)
		i++
	}
	pi = math.Pow(6*sumx, 0.5)
	fmt.Printf("%.02f \n", pi)
}
