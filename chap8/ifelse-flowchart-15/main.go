package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("ขอจำนวนที่ 1 (a) : ")
	fmt.Scanf("%f", &a)
	fmt.Print("ขอจำนวนที่ 2 (b) : ")
	fmt.Scanf("%f", &b)
	fmt.Print("ขอจำนวนที่ 3 (c) : ")
	fmt.Scanf("%f", &c)

	if a + b > c {
		fmt.Println("a + b > c")
	}
}
