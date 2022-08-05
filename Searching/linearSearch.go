package main

import "fmt"

func linSearch(arr []int, target int) (int, bool) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i, true
		}
	}
	return -1, false
}
func main() {
	var arr = []int{12, 43, 65, 23}
	Index, Result := linSearch(arr, 3)
	if Result || Index != -1 {
		fmt.Printf("Element found at index :%d", Index)
	} else {
		fmt.Printf("Element is not present in given array")
	}
}
