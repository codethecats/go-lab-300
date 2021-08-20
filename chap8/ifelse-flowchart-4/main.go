package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Print("Enter a number : ")
	fmt.Scanf("%f", &number)

	if number > 0 {
		if math.Mod(number, float64(2)) == 0 {
			fmt.Println("positive even")
		} else {
			fmt.Println("positive odd")
		}
	} else if number < 0 {
		if math.Mod(number, float64(2)) == 0 {
			fmt.Println("negative even")
		} else {
			fmt.Println("negative odd")
		}
	} else {
		fmt.Println("zero")
	}
}
