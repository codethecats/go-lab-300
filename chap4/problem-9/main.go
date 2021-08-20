package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Print("Enter a number: ")
	var a int
	fmt.Scanln(&a)

	fmt.Print("Enter a number: ")
	var b int
	fmt.Scanln(&b)

	var c int = a + b
	var d int = a - b
	var e int = a * b
	var f int = a / b

	g := reflect.TypeOf(a)
	h := reflect.TypeOf(b)
	i := reflect.TypeOf(c)
	j := reflect.TypeOf(d)
	k := reflect.TypeOf(e)
	l := reflect.TypeOf(f)

	fmt.Printf("%2d %s + %2d %s = %2d %s \n", a, g, b, h, c, i)
	fmt.Printf("%2d %s - %2d %s = %2d %s \n", a, g, b, h, d, j)
	fmt.Printf("%2d %s x %2d %s = %2d %s \n", a, g, b, h, e, k)
	fmt.Printf("%2d %s ÷ %2d %s = %2d %s \n", a, g, b, h, f, l)

}
