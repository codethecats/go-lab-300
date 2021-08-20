package main

import "fmt"

func main() {
	christianEra := 0

	fmt.Print("กรุณากรอกปีคริสตศักราช : ")
	fmt.Scanf("%d", &christianEra)

	if christianEra >= 0 {
		buddhistEra := christianEra + 543
		fmt.Printf("%s %d\n","พุทธศักราช =" ,buddhistEra)
	} else {
		fmt.Println("กรุณากรอกข้อมูลที่มากกว่าหรือเท่ากับ 0")
	}
}