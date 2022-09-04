package main

import "fmt"

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer
 */
func sol(arr []int, idx, sum, B int) int {
	// Base case
	if sum > 1000 {
		return 0
	}
	// Base case
	if idx == len(arr) {
		return 0
	}
	// Terminating condition
	if B == 0 {
		return 1
	}
	// Choose case + Not choose case
	return sol(arr, idx+1, sum+arr[idx], B-1) + sol(arr, idx+1, sum, B)
}
func solve(A []int, B int) int {
	return sol(A, 0, 0, B)
}

func main() {
	arr := []int{12, 32, 34}
	fmt.Println(solve(arr, 1))
}
