package main

import "fmt"

func main() {
	var num int

	fmt.Print("กรุณากรอกจำนวนเต็ม : ")
	fmt.Scanf("%d", &num)

	if num % 3 == 0 {
		fmt.Println("หารด้วย 3 ลงตัว")
	} else {
		fmt.Println("หารด้วย 3 ไม่ลงตัว")
	}
}
