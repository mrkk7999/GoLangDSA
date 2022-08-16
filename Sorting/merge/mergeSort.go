package main

import "fmt"

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
		} else {
			final = append(final, b[j])
		}
	}
	for i < len(a) {
		final = append(final, a[i])
	}

	for j < len(b) {
		final = append(final, b[j])
	}
	return final
}
func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	first := sort(arr[:len(arr)/2])
	second := sort(arr[len(arr)/2:])
	return merge(first, second)

}
func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9}
	fmt.Println(unsorted)
	sorted := sort(unsorted)
	fmt.Println(sorted)
}
