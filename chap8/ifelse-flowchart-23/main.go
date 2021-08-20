package main

import "fmt"

func main() {
	var listenTime float64

	fmt.Print("ใส่เวลา : ")
	fmt.Scanf("%f", &listenTime)

	if listenTime > 4 {
		fmt.Println("อัตรายต่อหู")
	} else {
		fmt.Println("ขอให้มีความสุขกับการฟังเพลง")
	}
}