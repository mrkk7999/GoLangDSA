package main

// Time complexity
// O(N^2) = O(N) + O(N-1) + ... + 1

func modulo(val int) int {
	return val % (10e+7)
}
func BruteForce(A []int) int {
	var invCount = 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			if A[i] > A[j] && i < j {
				invCount = modulo((modulo(invCount) + 1))
			}
		}
	}
	return modulo(invCount)
}

// using merge sort
// Time complexity - O(NlogN)
// Space complexity - O(N)

// we'll get inversion count by counting it when we'll start merging the divided arrays together
// if we've to put the element from right array into the resultant array we'll add the remaining elements
// on left into the inversion count

func Optimal() {

}
