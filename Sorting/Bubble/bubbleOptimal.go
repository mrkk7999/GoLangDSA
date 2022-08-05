package main

import "fmt"

func BubbleSortOptimal(arr *[]int) {
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			Swapped := false
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
				Swapped = true
			}
			if Swapped == false {
				break
			}
		}
	}
}

func main() {
	var arr = []int{12, 43, 65, 23}
	fmt.Println("Array before rotating", arr)
	BubbleSortOptimal(&arr)
	fmt.Println("Array after sorting in ascending order:", arr)
}
