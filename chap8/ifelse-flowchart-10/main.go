package main

import "fmt"

func main() {

	var num1 float64
	var num2 float64

	fmt.Print("กรุณากรอกจำนวนที่ 1 : ")
	fmt.Scanf("%f", &num1)

	fmt.Print("กรุณากรอกจำนวนที่ 2 : ")
	fmt.Scanf("%f", &num2)

	if num1 > num2 {
		fmt.Printf("%s %.1f\n","number =", num1)
	} else if num2 > num1 {
		fmt.Printf("%s %.1f\n","number2 =", num2)
	} else {
		fmt.Println("มีค่าเท่ากัน")
	}
}
