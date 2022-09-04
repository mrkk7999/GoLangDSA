package main

import "fmt"

func Facto(N int) int {
	if N == 0 {
		return 1
	}
	return N * Facto(N-1)
}

func main() {
	fmt.Println("Enter number to find factorial of:")
	var N int
	fmt.Scanln(&N)
	fmt.Printf("Entered Number %d factorial is %d", N, Facto(N))
}
