package main

import "fmt"

func main() {
	var num int

	fmt.Print("กรุณากรอกตัวเลขจำนวนเต็ม : ")
	fmt.Scanf("%d", &num)

	if num % 3 == 0 && num % 5 == 0 {
		fmt.Println("หารด้วย 3 และ 5 ลงตัว")
	} else {
		fmt.Println("หารด้วย 3 และ 5 ไม่ลงตัว")
	}
}