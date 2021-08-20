package main

import "fmt"

func main() {
	var bill, hour, minute, price int

	fmt.Print("ยอด shopping : ")
	fmt.Scanf("%d", &bill)
	fmt.Print("จำนวน ชั่วโมง ที่จอดรถ : ")
	fmt.Scanf("%d", &hour)
	fmt.Print("จำนวน นาที ที่จอดรถ : ")
	fmt.Scanf("%d", &minute)

	if minute > 0 {
		hour += 1
	}

	if bill >= 1000 {
		price = 30*(hour-4)
		if price < 0 {
			price = 0
		}
	} else {
		price = 30*(hour-1)
	}

	fmt.Println(price)
}
