package main

type treeNode struct {
	left  *treeNode
	Node  int
	right *treeNode
}

func newTreeNode(NodeVal int) *treeNode {
	var node *treeNode = new(treeNode)
	node.Node = NodeVal
	node.left = nil
	node.right = nil
	return node
}
