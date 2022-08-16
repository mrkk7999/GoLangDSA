package main

import "fmt"

// https://www.scaler.com/academy/mentee-dashboard/class/13343/assignment/problems/200/submissions
func binSearchItr(arr []int, target int, left, right int) (int, bool) {
	for left <= right {
		mid := left + (right-left)/2
		if target == arr[mid] {
			return mid, true
		} else if target < arr[mid] {
			right = mid - 1
		} else if target > arr[mid] {
			left = mid + 1
		}
	}
	return -1, false
}

func main() {
	var arr = []int{12, 43, 65, 23}
	Index, Result := binSearchItr(arr, 43, 0, len(arr)-1)
	if Result || Index != -1 {
		fmt.Printf("Element found at index :%d", Index)
	} else {
		fmt.Printf("Element is not present in given array")
	}
}
