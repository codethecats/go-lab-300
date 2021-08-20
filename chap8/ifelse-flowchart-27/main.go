package main

import "fmt"

func main() {
	var isMember string
	var bill, discount float64

	fmt.Print("isMember (yes/no) : ")
	fmt.Scanf("%s", &isMember)
	fmt.Print("ยอดซื้อ : ")
	fmt.Scanf("%f", &bill)

	if isMember == "yes" {
		if bill >= 5000 {
			discount = bill * 0.15
		} else if bill >= 1000 {
			discount = bill * 0.1
		} else if bill >= 500 {
			discount = bill * 0.05
		} else {
			discount = 0
		}
		fmt.Printf("%s %.1f\n","ส่วนลด =", discount)
	} else if isMember == "no" {
		discount = 0
		fmt.Printf("%s %.1f\n","ส่วนลด =", discount)
	} else {
		fmt.Println("กรุณากรอกข้อมูลให้ถูกต้อง")
	}
}
