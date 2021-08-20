package main

import "fmt"

func main() {
	var order string

	fmt.Print("สั่งอะไร (ไข่ดาว/ไข่เจียว/ไข่ต้ม) : ")
	fmt.Scanf("%s", &order)

	if order == "ไข่ดาว" {
		fmt.Println("แนะนำให้สั่งคู่กับต้มจืดไข่น้ำ")
	} else if order == "ไข่เจียว" {
		fmt.Println("แนะนำให้สั่งคู่กับไข่ลูกเขย")
	} else if order == "ไข่ต้ม" {
		fmt.Println("แนะนำให้สั่งคู่กับยำไข่ดาว")
	} else {
		fmt.Println("ร้านเรามี ไข่ดาว, ไข่เจียว, ไข่ต้ม")
	}
}
