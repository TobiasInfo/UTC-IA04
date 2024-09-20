package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	// Write the first hello world program in Go
	var h string = HelloWorld()
	fmt.Println(h)

	fmt.Println("Rayon : ")
	var r float64
	fmt.Scanf("%f", &r)

	out := computeRayon(r)
	fmt.Println(out)

	forExample()
}
