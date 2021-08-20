package main

import "fmt"

func main() {
	var velocity float64

	fmt.Print("กรุณากรอกความเร็ว (km/hr) : ")
	fmt.Scanf("%f", &velocity)

	if velocity > 120 {
		fmt.Println("ออกใบสั่ง")
	}
}
