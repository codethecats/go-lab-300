package main

import (
	"fmt"
)

func main() {
	var hour, minute int

	fmt.Print("จำนวน ชั่วโมง ที่จอดรถ : ")
	fmt.Scanf("%d", &hour)
	fmt.Print("จำนวน นาที ที่จอดรถ : ")
	fmt.Scanf("%d", &minute)

	if minute > 0 && minute < 60 {
		hour += 1
	}

	price := 30*(hour-1)

	fmt.Printf("%s %d %s\n", "ค่าจอดรถ :", price, "บาท")
}
