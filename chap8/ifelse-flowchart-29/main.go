package main

import "fmt"

func main() {
	var order string
	var amount, cost int

	fmt.Print("สั่งอะไร (ไข่ดาว/ไข่เจียว/ไข่ต้ม) : ")
	fmt.Scanf("%s", &order)
	fmt.Print("จำนวนที่สั่ง : ")
	fmt.Scanf("%d", &amount)

	if order == "ไข่ดาว" {
		cost = amount * 7
		fmt.Printf("%s %d\n","ราคา =", cost)
	} else if order == "ไข่เจียว" {
		cost = amount * 10
		fmt.Printf("%s %d\n","ราคา =", cost)
	} else if order == "ไข่ต้ม" {
		cost = amount * 5
		fmt.Printf("%s %d\n","ราคา =", cost)
	} else {
		fmt.Println("ร้านเรามี ไข่ดาว, ไข่เจียว, ไข่ต้ม")
	}
}
