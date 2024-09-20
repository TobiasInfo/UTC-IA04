package main

import (
	"fmt"
	"sync"
)

var (
	n  = 0
	mu sync.Mutex
	wg sync.WaitGroup
)

func f() {
	defer wg.Done()
	mu.Lock()
	n++
	mu.Unlock()
}

func main() {
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go f()
	}
	wg.Wait()

	fmt.Println("n:", n)
}
