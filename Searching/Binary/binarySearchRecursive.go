package main

import "fmt"

func binSearchRec(arr []int, target int, left, right int) (int, bool) {
	for left <= right {
		mid := left + (right-left)/2
		if target == arr[mid] {
			return mid, true
		} else if target < arr[mid] {
			binSearchRec(arr, target, left, mid-1)
		} else if target > arr[mid] {
			binSearchRec(arr, target, mid+1, right)
		}
	}
	return -1, false
}

func main() {
	var arr = []int{12, 43, 65, 23}
	Index, Result := binSearchRec(arr, 43, 0, len(arr)-1)
	if Result || Index != -1 {
		fmt.Printf("Element found at index :%d", Index)
	} else {
		fmt.Printf("Element is not present in given array")
	}
}
