package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var empty = Node{
	Value: nil,
	Next:  nil,
}
var sizeOfLinkedList = 0

func NewNode(Value int) *Node {
	var NewNode = new(Node)
	NewNode.Next = nil
	NewNode.Value = Value
	return NewNode
}

func PrintLinkedList(Head *Node) {
	var temp = Head
	var flag = 0
	fmt.Println()
	for temp != nil {
		if flag == 0 {
			fmt.Print(temp.Value)
			flag = 1
		} else {
			fmt.Print(" ", temp.Value)
		}
		temp = temp.Next
	}
	fmt.Println()
}

func AddNodeAtPosition(Head *Node, Position, NodeValue int) {
	var prev = Head
	var count = 1
	var temp = NewNode(NodeValue)
	if Head == nil {
		Head = temp
	}
	if Position >= 1 && Position < sizeOfLinkedList {
		for count <= Position+1 {
			prev = prev.Next
		}
	}
	if Position == sizeOfLinkedList {

	}
	// prev = prev.Next
	temp.Next = prev.Next
	prev.Next = temp
	fmt.Println()
	PrintLinkedList(Head)
	fmt.Println()
	sizeOfLinkedList++
}

func DeleteNodeAtPosition(Head *Node, Position int) {
	var prev = Head
	fmt.Println()
	PrintLinkedList(Head)
	fmt.Println()
	PrintLinkedList(prev)
	fmt.Println()
	if Position < 1 || Position > sizeOfLinkedList {
		return
	} else {

	}
	//var count = 0
	//if Position >= 1 && Position <= SizeOfLinkedList(Head) {
	//	for count < Position {
	//		prev = prev.Next
	//	}
	//	count++
	//}
	//if Position == SizeOfLinkedList(Head) {
	//	prev.Next = nil
	//} else {
	//	prev.Next = prev.Next.Next
	//}
}

func InsertNodeAtFront(Head *Node, Value int) {
	var temp = NewNode(Value)
	temp.Next = Head
	Head = temp
}

func InsertNodeAtLast(Head *Node, Value int) {
	var temp = Head
	var lastNode = NewNode(Value)
	for temp != nil {
		temp = temp.Next
	}
	temp.Next = lastNode
}

func DeleteNodeAtFirst() {

}

func DeleteNodeAtLast() {

}

func CopyLinkedList() {

}

func RemoveDuplicates() {

}

func SizeOfLinkedList(Head *Node) int {
	// tempNode for iteration
	var tempNode = Head
	var count = 0
	for tempNode != nil {
		count = count + 1
		tempNode = tempNode.Next
	}
	return count
}

func main() {
	var Head = NewNode(11)
	sizeOfLinkedList++
	//Head.Next = nil
	//Head.Value = 11
	//sizeOfLinkedList++
	AddNodeAtPosition(Head, 1, 111)
	PrintLinkedList(Head)
	fmt.Println(sizeOfLinkedList)
	AddNodeAtPosition(Head, 2, 12)
	PrintLinkedList(Head)
	fmt.Println(sizeOfLinkedList)
	AddNodeAtPosition(Head, 3, 13)
	fmt.Println(sizeOfLinkedList)
	PrintLinkedList(Head)
	AddNodeAtPosition(Head, 4, 14)
	AddNodeAtPosition(Head, 5, 15)
	PrintLinkedList(Head)
	AddNodeAtPosition(Head, 6, 16)
	PrintLinkedList(Head)
	DeleteNodeAtPosition(Head, 4)
	PrintLinkedList(Head)
}
