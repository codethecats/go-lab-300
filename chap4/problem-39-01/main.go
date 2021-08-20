package main

import (
	"bytes"
	"fmt"
)

func main() { /*
	วิธีต่อสตริง Golang
	string concatenation เป็นพื้นฐานที่สุดในการเขียนโปรแกรมหลาย ๆ ภาษาจะใช้เครื่องหมาย + ในภาษา Go ก็เช่นกัน
	มาดูกันครับ มีแบบต่างๆให้เราเลือกใช้แบบไหนบ้าง

	ใช้งานฟังก์ชัน bytes.WriteString ในการต่อสตริง และแปลงให้อยู่ในรูปแบบตัวแปรสตริงด้วยฟังก์ชัน bytes.String
*/

	// 1.ใช้งาน buffer Package byte
	str1 := ""
	str2 := ""

	fmt.Print("Please insert str1: ")
	fmt.Scanf("%s", &str1)

	fmt.Print("Please insert str2: ")
	fmt.Scanf("%s", &str2)

	var str bytes.Buffer

	// เริ่มต้นด้วย "Hello "
	str.WriteString(str1)

	// ต่อคำว่า world! ลงในตัวแปร
	str.WriteString(str2)

	fmt.Println(str.String()) // Hello world!
}
