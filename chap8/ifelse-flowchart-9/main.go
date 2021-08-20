package main

import "fmt"

func main() {
	var thb float64

	fmt.Print("กรุณากรอกจำนวนเงินในหน่วยบาท : ")
	fmt.Scanf("%f", &thb)

	if thb > 0 {
		usd := thb / 32.8
		bankProfit := 0.3 * usd

		fmt.Printf("%s %.1f %s\n","จำนวนเงินในหน่วยดอลลาร์ =", usd, "USD")
		fmt.Printf("%s %.1f %s\n","ธนาคารจะได้กำไร =", bankProfit, "บาท")

	} else {
		fmt.Println("you don't have money")
	}
}
