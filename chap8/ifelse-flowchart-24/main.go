package main

import "fmt"

func main() {
	var status string
	var distance, fee float64

	fmt.Print("กรุณากรอกสภาพการจราจร : ")
	fmt.Scanf("%s", &status)

	fmt.Print("ระยะทาง : ")
	fmt.Scanf("%f", &distance)

	if status == "คล่องตัว" {
		fee = distance * 10
		fmt.Printf("%s %.1f\n","ค่าเดินทาง", fee)
	} else if status == "ปานกลาง" {
		fee = distance * 12
		fmt.Printf("%s %.1f\n","ค่าเดินทาง", fee)
	} else if status == "หนาแน่น" {
		fee = distance * 15
		fmt.Printf("%s %.1f\n","ค่าเดินทาง", fee)
	}
}
