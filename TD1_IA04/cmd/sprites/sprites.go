package main

import (
	f "Td1_IA04/go/pkg/formes"
	"fmt"
)


func main() {
	p := f.NewPoint2D(3, 4)
	fmt.Println("Initial point:", p)

	fmt.Println("X:", p.X())
	fmt.Println("Y:", p.Y())

	p.SetX(5)
	p.SetY(12)
	fmt.Println("Point after setting X and Y:", p)

	clonedPoint := p.Clone()
	fmt.Println("Cloned point:", clonedPoint)

	fmt.Println("Module of the point:", p.Module())
}

