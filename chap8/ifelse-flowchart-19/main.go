package main

import "fmt"

func main() {
	var bill, discount, pay float64

	fmt.Print("กรุณากรอกยอดซื้อ : ")
	fmt.Scanf("%f", &bill)

	if bill >= 50000 {
		discount = bill * 0.2
	} else if bill >= 10000 {
		discount = bill * 0.15
	} else if bill >= 1000 {
		discount = bill * 0.1
	} else {
		discount = 0
	}

	pay = bill - discount
	fmt.Printf("%s %.1f\n","ยอดเงินที่ลูกค้าต้องจ่าย =", pay)
}
