package main

// https://www.scaler.com/academy/mentee-dashboard/class/9578/homework/problems/4073/submissions
// https://www.geeksforgeeks.org/find-elements-array-least-two-greater-elements/

import (
	"fmt"
	"math"
)

func solve(A []int) []int {
	var largest, seclargest = math.MinInt, math.MinInt
	var res []int
	for i := 0; i < len(A); i++ {
		if A[i] > largest {
			seclargest = largest
			largest = A[i]
		} else if A[i] > seclargest {
			seclargest = A[i]
		}
	}
	for i := 0; i < len(A); i++ {
		if A[i] < seclargest {
			res = append(res, A[i])
		}
	}

	return res
}

func main() {
	var arr = []int{21, 45, 23, 87, 12}
	fmt.Println(solve(arr))
}
