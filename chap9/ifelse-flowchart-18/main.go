package main

import "fmt"

func main() {
    for i := 0; i <= 7; i++ {
        for j := 6; j >= -2; j -= 1 {
            fmt.Println(i, j)
        }
    }
}
