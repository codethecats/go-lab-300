package main

import "fmt"

func main() {
	var number float64

	fmt.Print("Enter a number : ")
	fmt.Scanf("%f", &number)

	if number > 0{
		fmt.Printf("Positive \n")
	} else if number < 0{
		fmt.Printf("Negative \n")
	} else {
		fmt.Printf("Zero \n")
	}
}
