package main

import "fmt"

func main() {
    for i := -6; i <= 7; i += 2 {
        for j := 1; j <= 8; j += 3 {
            fmt.Println(i, j)
        }
    }
}
