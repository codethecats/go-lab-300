package main

import "fmt"

func main() {
    for i := 12; i >= 0; i -= 2{
        for j := -5; j <= 0; j++ {
            fmt.Println(i, j)
        }
    }
}
