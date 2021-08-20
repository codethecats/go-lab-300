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
	var b float64
	fmt.Scanln(&b)

	var c int = int(float64(a) + b)
	var d int = int(float64(a) - b)
	var e int = int(float64(a) * b)
	var f int = int(float64(a) / b)

	g := reflect.TypeOf(a)
	h := reflect.TypeOf(b)
	i := reflect.TypeOf(c)
	j := reflect.TypeOf(d)
	k := reflect.TypeOf(e)
	l := reflect.TypeOf(f)

	fmt.Printf("%2d %s + %2.2f %s = %2d %s \n", a, g, b, h, c, i)
	fmt.Printf("%2d %s - %2.2f %s = %2d %s \n", a, g, b, h, d, j)
	fmt.Printf("%2d %s x %2.2f %s = %2d %s \n", a, g, b, h, e, k)
	fmt.Printf("%2d %s ÷ %2.2f %s = %2d %s \n", a, g, b, h, f, l)

}
