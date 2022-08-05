package main

import "fmt"

func singleNonDuplicate(nums []int) (int, int) {
	N := len(nums)
	left, right := 0, N-1
	for left <= right {
		M := left + (right-left)/2
		if (M == 0 || nums[M] != nums[M-1]) && (M == N-1 || nums[M] != nums[M+1]) {
			return nums[M], M
		}
		if M == 0 || nums[M] != nums[M-1] {
			if M%2 == 0 {
				left = M + 1
			} else {
				right = M - 1
			}
		} else {
			if M%2 == 0 {
				right = M - 1
			} else {
				left = M + 1
			}
		}
	}
	return -1, -1
}

func main() {
	var arr = []int{12, 42, 42, 43, 43}
	element, index := singleNonDuplicate(arr)
	if index != -1 {
		fmt.Printf("Single element found at index: %d which is : %d", index, element)
	} else {
		fmt.Println("No single element")
	}
}
