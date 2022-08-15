package main

import (
	"fmt"
	"math"
)

func kthsmallest(A []int, B int) int {
	for i := 0; i < B; i++ {
		min := math.MaxInt
		minIdx := i
		for j := i; j < len(A); j++ {
			if A[j] < min {
				min = A[j]
				minIdx = j
			}
		}
		A[i], A[minIdx] = min, A[i]
	}
	return A[B-1]
}

func main() {
	var arr = []int{12, 32, 1, 54, 2}
	fmt.Println(kthsmallest(arr, 3))
}
