package main

import "fmt"

func main() {
	width := 0
	length := 0

	fmt.Print("Enter a width : ")
	fmt.Scanf("%d", &width)
	fmt.Print("Enter a length : ")
	fmt.Scanf("%d", &length)

	if width > 0 && length > 0 {
		area := width * length
		fmt.Println(area)
	} else {
		fmt.Println("กรุณากรอกจำนวนเต็มบวก")
	}
}