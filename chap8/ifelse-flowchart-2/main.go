package main

import "fmt"

func main() {
	number := 0

	fmt.Print("Enter a number : ")
	fmt.Scanf("%d", &number)

	if number > 0 {
		fmt.Printf("มากกว่า 0 \n")
	} else {
		fmt.Printf("น้อยกว่าหรือเท่ากัย 0 \n")
	}
}
