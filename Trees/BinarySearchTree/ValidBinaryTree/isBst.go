package main

import "fmt"

// https://www.scaler.com/academy/mentee-dashboard/class/15706/assignment/problems/221
// https://leetcode.com/problems/validate-binary-search-tree/

type treeNode struct {
	left  *treeNode
	Node  int
	right *treeNode
}

func NewTreeNode(NodeValue int) *treeNode {
	var node = new(treeNode)
	node.Node = NodeValue
	node.left = nil
	node.right = nil
	return node
}

func validate(root *treeNode, max, min *int) bool {
	if root == nil {
		return true
	}
	if (max != nil && root.Node >= *max) || (min != nil && root.Node <= *min) {
		return false
	}

	return validate(root.left, &(root.Node), min) && validate(root.right, max, &(root.Node))
}
func main() {
	//min := int(2e9)
	//max := int(-2e9)
	var tree = new(treeNode)
	tree.Node = 5
	tree.left.Node = 4
	tree.left.left.Node = 3
	tree.left.right.Node = 6
	tree.right.Node = 8
	tree.right.left.Node = 7
	tree.right.right.Node = 9
	fmt.Println(validate(tree, nil, nil))
}
