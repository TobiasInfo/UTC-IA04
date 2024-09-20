package main

import (
    "fmt"
)

var n = 0

func f() {
    n++
}

func main() {
    for i := 0; i < 10000; i++ {
        go f()
    }

    fmt.Println("Appuyez sur entrée")
    fmt.Scanln()
    fmt.Println("n:", n)
}
