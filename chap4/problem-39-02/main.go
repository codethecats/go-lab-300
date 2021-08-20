package main

import "fmt"

func main() {
	// 2. ใช้งานฟังก์ชัน copy
	str1 := ""
	str2 := ""

	fmt.Print("Please insert str1: ")
	fmt.Scanf("%s", &str1)

	fmt.Print("Please insert str2: ")
	fmt.Scanf("%s", &str2)

	str := make([]byte, 13)
	i := 0

	// เริ่มต้นด้วย "Hello "
	i += copy(str[i:], str1)

	// ต่อคำว่า world! ต่อลงไปใน byte
	i += copy(str[i:], str2)
	fmt.Println(string(str)) // Hello, world!
}
