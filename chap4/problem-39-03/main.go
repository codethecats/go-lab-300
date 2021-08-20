package main

import (
	"fmt"
	"strings"
)

func main() {
	// ตัวแปร slice ของ byte สามารถ convert type เป็นสตริงได้เลย
	// 3. ใช้งานฟังก์ชัน join
	str1 := ""
	str2 := ""

	fmt.Print("Please insert str1: ")
	fmt.Scanf("%s", &str1)

	fmt.Print("Please insert str2: ")
	fmt.Scanf("%s", &str2)

	str := make([]string, 0)
	str = append(str1, str2)

	// ต่อคำว่า world! ต่อลงไปใน string
	// str = append(str, "world!")
	fmt.Println(strings.Join(str, "")) // Hello, world!
}
