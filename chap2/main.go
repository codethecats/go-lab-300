package main

import (
	"fmt"
)

func main() {

//1
fmt.Print("Hello")

//2
fmt.Print("Hello World")

//3
fmt.Print("hello_golang")

//4
fmt.Print("Merry X'Mas")

//5
fmt.Print(`I want to ask you "Why don't you drive to work ?"`)

//6
fmt.Print(`I don't have a car`)

//7
fmt.Print(`You got a new job !? That's so exciting`)

//8
fmt.Print("สวัสดีวันจันทร์")

//9
fmt.Print(`ความแตกต่างระหว่างคนเก่งกับคนไม่เก่ง คือ "การใช้เวลาว่างให้เป็นประโยชน์"`)

//10
fmt.Print(` /\\/\\/\\ `)

//11
fmt.Println(`a
an
ant`)

//12
fmt.Println(`   *
*  *  *
   *`)

//13
fmt.Println(`*	 +	 *
+	*	+
*	+	*`)

//14
fmt.Println(`Just because something
thinks differently from you,
dose that mean it's not thinking ?`)

//15
fmt.Println(`
\		/
 	x
/		\`)

//16
fmt.Println(25)

//17
absoluteZero := 100.000000
fmt.Printf("%.6f \n", absoluteZero)

//18
fmt.Println(3.141592653589793)

//19
a := 2
fmt.Println(a)

//20
a := 12.5
fmt.Println(a)

//21
a := 2
b := 3

fmt.Printf("%d x %d = %d \n", a, b, a*b)

//22
a := 2
b := 3

fmt.Printf("%d + %d = %d + %d = %d \n", a, b, b, a, a+b)

//23
a := 2
b := 3
c := 5

fmt.Printf("%d*(%d+%d) = %d*%d + %d*%d \n", a, b, c, a, b, a, c)

//24
var a float64 = 2.4
var b float64 = 2.5

fmt.Printf("%f + %f = %.4f \n", a, b, a+b)

//25
var a, b int = 5, 2
d := float64(int(a))
e := float64(int(b))
var c float64 = float64(int(a - b))
fmt.Printf("%.1f - %.2f = %.4f \n", d, e, c)

//26
birthday := 25
fmt.Println("ฉันเกิดวันที่",birthday,"ธันวาคม")

//27
a := 5
b := 100
fmt.Printf(a," %d เท่าของ",b,"มีค่าเท่่ากับ","500")
fmt.Printf("%d เท่าของ %d มีค่าเท่่ากับ %d \n", a, b, a*b)

//28
a := 3.5
fmt.Printf("เขามีเงินเยอะกว่าฉัน %.2f บาท \n", a)

//29
a := 5
fmt.Println("ฉันได้กำไร",a,"%")

//30
a := 2
b := 3.5
fmt.Println("เมื่อวานฉันขาดทุน",a,"%","วันนี้ฉันได้กำไร",b,"%")

}