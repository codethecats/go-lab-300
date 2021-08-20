package main

import "fmt"

func main() {
	list1 := []interface{}{}

	str1 := ""
	str2 := ""
	str3 := ""

	fmt.Print("Please insert str1: ")
	fmt.Scanf("%s", &str1)
	fmt.Print("Please insert str2: ")
	fmt.Scanf("%s", &str2)
	fmt.Print("Please insert str3: ")
	fmt.Scanf("%s", &str3)

	list1 = append(list1, str1)
	list1 = append(list1, str2)
	list1 = append(list1, str3)

	fmt.Printf("%#v\n", list1)
}
