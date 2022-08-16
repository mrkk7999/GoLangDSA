package main

import "fmt"

/**
 * @input A : Integer
 *
 * @Output Integer
 */
func sqrt(A int) int {
	// return getSqrtOfA(A)
	// taking int64 cause calculation might need higher integer
	var low, high, mid, ans int64 = 1, int64(A), 0, 0
	// breaking condition
	for low <= high {
		mid = low + (high-low)/2
		// as we've to found floor value here, so we're checking for value matching or
		// nearest lower not nearesthigher
		if (mid * mid) <= int64(A) {
			// answer lies in second half
			low = mid + 1
			ans = mid
		} else {
			// answer lies in first half
			high = mid - 1
		}
	}
	return int(ans)
}

func main() {
	fmt.Println("Enter value you want to find sqaure root of")
	var x int
	fmt.Scan(&x)
	fmt.Println(sqrt(x))
}
