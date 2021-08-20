package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("กรุณากรอกอุณหภูมิองศาฟาเรนไฮต์ : ")
	fmt.Scanf("%f", &fahrenheit)

	if fahrenheit >= 32 {
		celsius := (fahrenheit - 32) * 5.0/9.0
		fmt.Printf("%s %.2f\n","celsius =", celsius)
	} else {
		fmt.Printf("cold\n")
	}
}
