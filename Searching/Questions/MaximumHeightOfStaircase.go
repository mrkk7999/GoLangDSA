// Questions
// Given an integer A representing the number of square blocks. The height of each square block is 1. The task is to create a staircase of max-height using these blocks.

// The first stair would require only one block, and the second stair would require two blocks, and so on.

// Find and return the maximum height of the staircase.

package main

import "fmt"

// Problem approach
// keep two variables sum and count 
// keep adding to sum and checking if it's not crossing the given integer value
// if it's then break the loop and come out of the loop and return the count
// else increase the count counter

func solve(A int) int {
	var sum, count = 0, 0
	for i := 1; i <= A; i++ {
		// make sure you'll add first and then check the breaking condition
		sum = sum + i
		if sum > A {
			break
		}
		count = count + 1
	}
	return count
}

func main() {
	var intVar = 12
	fmt.Println(solve(intVar))
}
