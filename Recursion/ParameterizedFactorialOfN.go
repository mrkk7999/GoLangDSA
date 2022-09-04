package main

import "fmt"

func Fact(N, fact int) int {
	if N == 0 {
		fmt.Println(fact)
		return fact
	}
	return Fact(N-1, fact*N)
}

func main() {
	fmt.Println("Enter number to find factorial of:")
	var N int
	fmt.Scanln(&N)
	fmt.Printf("Entered Number %d factorial is %d", N, Fact(N, 1))
}
