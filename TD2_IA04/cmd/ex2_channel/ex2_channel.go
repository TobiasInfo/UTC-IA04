package main

import (
	"fmt"
)

func main() {
	n := 0
	ch := make(chan int) 

	f := func() {
		ch <- 1 // Envoie un signal de 1 pour signifier une incrémentation
	}

	for i := 0; i < 10000; i++ {
		go f()
	}

	// Recevoir les signaux des goroutines et incrémenter n
	for i := 0; i < 10000; i++ {
		n += <-ch // Incrémenter n à chaque réception d'un signal
	}

	fmt.Println("n:", n)
}
