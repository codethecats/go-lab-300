package main

import "fmt"

func main() {
    for i := 14; i >= 10; i -= 1 {
        for j := 4; j >= -1; j -= 1 {
            fmt.Println(i, j)
        }
    }
}
