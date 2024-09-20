package main

// Write a function that return a string "Hello, World!"
// This function will be called from the main function

func HelloWorld() string {
	return "Hello, World!"
}

func computeRayon(r float64) float64 {

	const pi float64 = 3.141592653589793
	return 2 * pi * r
}

func forExample() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}
