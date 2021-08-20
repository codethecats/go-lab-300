package main

import "fmt"

func main() {
	var currency string
	var amount, thb float64

	fmt.Print("สกุลเงิน : ")
	fmt.Scanf("%s", &currency)
	fmt.Print("กรอกจำนวนเงิน : ")
	fmt.Scanf("%f", &amount)

	if currency == "USD" || currency == "usd"{
		thb = amount * 32.5
		fmt.Printf("%s %.1f\n","จำนวนเงินบาทที่แลกได้ =", thb)
	} else if currency == "JPY" || currency == "jpy"{
		thb = amount * 0.29
		fmt.Printf("%s %.1f\n","จำนวนเงินบาทที่แลกได้ =", thb)
	} else {
		fmt.Println("กรุณากรอกข้อมูลให้ถูกต้อง")
	}
}
