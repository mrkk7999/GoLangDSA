package main

import "fmt"

func Swap(arr *[]int, i, j int) {
	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp
}

func BubbleSort(arr *[]int) {
	length := len(*arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {

			if (*arr)[j] > (*arr)[j+1] {
				Swap(arr, j, j+1)
				//temp := (*arr)[j]
				//(*arr)[j] = (*arr)[j+1]
				//(*arr)[j+1] = temp
			}
		}
	}
}

func main() {
	var arr = []int{12, 43, 65, 23}
	fmt.Println("Array before rotating", arr)
	BubbleSort(&arr)
	fmt.Println("Array after sorting in ascending order:", arr)
}
