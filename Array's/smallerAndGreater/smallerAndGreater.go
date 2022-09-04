package main

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
import (
	// "sort"
	// "fmt"
	"fmt"
	"math"
)

func solve(A []int) int {
	// sort.Ints(A)
	// if len(A) <= 2 {
	//     return 0
	// }
	// return len(A)-2
	// Time complexity - O(N)
	var min, max, count = math.MaxInt, math.MinInt, 0
	for i := 0; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	for i := 0; i < len(A); i++ {
		if A[i] > min && A[i] < max {
			count = count + 1
		}
	}

	return count
}

func main() {
	var arr = []int{21, 45, 23, 87, 12}
	fmt.Println(solve(arr))
}
