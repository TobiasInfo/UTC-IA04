package main

import (
	"fmt"
	"time"
)

func goodYear() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Bonne année !")
}

func goodYear2() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		<-time.After(1 * time.Second)
	}
	fmt.Println("Bonne année !")
}

func goodYear3() {
	ticker := time.Tick(1 * time.Second)

	for i := 5; i > 0; i-- {
		fmt.Println(i)
		<-ticker
	}
	fmt.Println("Bonne année !")
}

func main() {
	goodYear()
	goodYear2()
	goodYear3()
}
