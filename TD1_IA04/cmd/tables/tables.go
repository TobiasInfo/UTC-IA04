package main

import (
	"Td1_IA04/go/pkg/fill"
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Tables")
	x := make([]int, 11)
	x = fill.Fill(x)
	fmt.Println(x)
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	newTable := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(newTable)
	fmt.Println(x)
	avg := fill.Moyenne(x)
	fmt.Println(avg)
	fmt.Println(fill.ValeursCentrales(x))
	fmt.Println(fill.Moyenne(fill.Plus1(x)))
}
