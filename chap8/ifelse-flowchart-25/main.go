package main

import "fmt"

func main() {
	var score int

	fmt.Print("กรุณากรอกคะแนน : ")
	fmt.Scanf("%d", &score)

	if score >= 50 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}
