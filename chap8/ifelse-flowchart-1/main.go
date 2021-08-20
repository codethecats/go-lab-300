package main

import "fmt"

func main() {
	number := 0
	fmt.Print("Please insert number: ")
	fmt.Scanf("%d", &number)

	if number%2 == 0 {
		fmt.Printf("เลขคู่ \n")
	} else {
		fmt.Printf("เลขคี่ \n")
	}
}
