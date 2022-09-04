package main

import "fmt"

type treeNode struct {
	left  *treeNode
	value int
	right *treeNode
	next  *treeNode
}

func helper(nums []int, low int, high int) *treeNode {
	if low < high {
		return nil
	}
	mid := (low + high) / 2
	var node = new(treeNode)
	node.value = mid
	node.left = helper(nums, low, mid-1)
	node.right = helper(nums, mid+1, high)
	return node
}
func sortedArrayToBST(nums []int) *treeNode {
	if len(nums) == 0 {
		return nil
	}
	return helper(nums, 0, len(nums)-1)
}

func main() {
	var arr = []int{21, 34, 23, 65, 32}
	fmt.Println(sortedArrayToBST(arr))
}
