package main

import "fmt"

func main() {
	var usd float64

	fmt.Print("กรุณากรอกจำนวนเงินในหน่วยดอลลาร์สหรัฐฯ : ")
	fmt.Scanf("%f", &usd)

	if usd > 0 {
		thb := usd*32.5
		fmt.Printf("%s %.2f\n","จำนวนเงินในหน่วยบาท =", thb)
	} else {
		fmt.Println("you don't have money")
	}
}
