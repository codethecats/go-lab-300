package main

import "fmt"

func main() {
	var character1 string
	var character2 string

	fmt.Print("กรุณากรอกสายอักขระที่ 1 : ")
	fmt.Scanf("%s", &character1)
	fmt.Print("กรุณากรอกสายอักขระที่ 2 : ")
	fmt.Scanf("%s", &character2)

	length1 := len(character1)
	length2 := len(character2)

	if length1 == length2 {
		fmt.Println("same")
	} else {
		fmt.Println("not same")
	}
}
