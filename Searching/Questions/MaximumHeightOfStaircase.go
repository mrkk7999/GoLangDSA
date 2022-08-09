package main

import "fmt"

func solve(A int) int {
	var sum, count = 0, 0
	for i := 1; i <= A; i++ {
		sum = sum + i
		if sum > A {
			break
		}
		count = count + 1
	}
	return count
}

func main() {
	var intVar = 12
	fmt.Println(solve(intVar))
}
