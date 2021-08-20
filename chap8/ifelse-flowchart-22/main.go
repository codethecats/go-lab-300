package main

import "fmt"

func main() {
	var velocity float64

	fmt.Print("กรุณากรอกความเร็ว (km/hr) : ")
	fmt.Scanf("%f", &velocity)

	if velocity > 120 {
		fmt.Println("ปรับ 1,000 บาท")
	} else if velocity > 80 {
		fmt.Println("ปรับ 500 บาท")
	}
}