// Problem Link
// https://www.scaler.com/academy/mentee-dashboard/class/13343/assignment/problems/4131/?navref=cl_pb_nv_tb
// https://leetcode.com/problems/single-element-in-a-sorted-array/

// Question
// You are given a sorted array consisting of only integers where every element appears exactly twice,
// except for one element which appears exactly once.
//Return the single element that appears only once.
// Your solution must run in O(log n) time and O(1) space.

package main

import "fmt"

// Approach
func singleNonDuplicate(nums []int) (int, int) {
	N := len(nums)
	left, right := 0, N-1
	// loop
	for left <= right {
		M := left + (right-left)/2
		// checks if single element present at the edges
		if (M == 0 || nums[M] != nums[M-1]) && (M == N-1 || nums[M] != nums[M+1]) {
			return nums[M], M
		}
		// it implies element at M & M+1 position are same
		if M == 0 || nums[M] != nums[M-1] {
			// that means our mid is at first position since in above condition we found mid and mid next are same and
			// if mid is even that means no single inserted till now
			// so we can go on right
			if M%2 == 0 {
				left = M + 1
			} else {
				// else we're going right
				right = M - 1
			}
		} else {
			// this means element at M & M's previous are same
			if M%2 == 0 {
				// if second element is even that means single element occured earlier
				// so go left
				// that means lower right
				right = M - 1
			} else {
				// go right
				// means increase left
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
