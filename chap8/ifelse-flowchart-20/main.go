package main

import "fmt"

func main() {
	var gender string
	var weight float64
	var height int

	fmt.Print("กรุณากรอกเพศ : ")
	fmt.Scanf("%s", &gender)
	fmt.Print("กรุณากรอกน้ำหนัก (kg) : ")
	fmt.Scanf("%f", &weight)
	fmt.Print("กรุณากรอกส่วนสูง (cm) : ")
	fmt.Scanf("%d", &height)

	if gender == "male" {
		if weight > float64(height - 100) {
			fmt.Println("ควรออกกำลังกาย")
		} else {
			fmt.Println("คุณผู้ชายหุ่นดีเยี่ยม")
		}
	} else if gender == "female" {
		if weight > float64(height - 110) {
			fmt.Println("ควรออกกำลังกาย")
		} else {
			fmt.Println("คุณผู้หญิงหุ่นดีเยี่ยม")
		}
	} else {
		fmt.Println("กรุณากรอกข้อมูลให้ถูกต้อง")
	}
}
