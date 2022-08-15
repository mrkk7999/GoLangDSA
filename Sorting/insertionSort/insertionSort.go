package main

import (
	"fmt"
	"math"
)

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var min, minIdx = math.MaxInt, i
		var unchanged = true
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minIdx = j
				unchanged = false
			}
		}
		if unchanged {
			return arr
		}
		arr[i], arr[minIdx] = min, arr[i]

	}
	return arr

}

func main() {
	var arr = []int{12, 32, 1, 54, 2}
	//fmt.Println(selectionSort(arr, true))

	fmt.Println(selectionSort(arr))
}
