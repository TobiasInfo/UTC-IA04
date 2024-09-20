package main

import (
    "fmt"
    "time"
)

func compte(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(i)
    }
}

func compteMsg(n int, msg string) {
	for i := 0; i < n; i++ {
		fmt.Println(msg, i)
	}
}

func compteMsgFromTo(start int, end int, msg string) {
    for i := start; i < end; i++ {
        fmt.Println(msg, i)
    }
}

func main() {
    fmt.Println("Appel normal de la fonction compte:")
    compte(5)

    // fmt.Println("\nAppel de la fonction compte avec une goroutine:")
    // go compte(5)

	// go compteMsg(5, "goroutine 1:")
	// go compteMsg(5, "goroutine 2:")

    for i := 0; i < 10; i++ {
        start := i * 10
        end := start + 10
        go compteMsgFromTo(start, end, fmt.Sprintf("Goroutine %d:", i+1))
    }
	
    time.Sleep(time.Second)
}
