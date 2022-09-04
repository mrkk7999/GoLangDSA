package main

//
//import (
//	"fmt"
//)
//
//type Node struct {
//	data int
//	next *Node
//}
//
//var root *Node = nil
//var sizeOfLL int = 0
//
//func insert_node(position, value int) {
//	// @params position, integer
//	// @params value, integer
//	if position >= 1 && position <= sizeOfLL+1 {
//		temp := new(Node)
//		temp.data = value
//		if position == 1 {
//			temp.next = root
//			root = temp
//		} else {
//			count := 1
//			prev := root
//			for count < position-1 {
//				prev = prev.next
//				count++
//			}
//			temp.next = prev.next
//			prev.next = temp
//		}
//		sizeOfLL++
//	}
//}
//
//func delete_node(position int) {
//	// @params position, integer
//	if position >= 1 && position <= sizeOfLL {
//		if position == 1 {
//			root = root.next
//		} else {
//			count := 1
//			prev := root
//			for count < position-1 {
//				prev = prev.next
//				count++
//			}
//			prev.next = prev.next.next
//		}
//		sizeOfLL--
//	}
//}
//
//func print_ll() {
//	// Output each element followed by a space
//	temp := root
//	flag := 0
//	for temp != nil {
//		if flag == 0 {
//			fmt.Print(temp.data)
//			flag = 1
//		} else {
//			fmt.Print(" ", temp.data)
//		}
//		temp = temp.next
//	}
//}
//
//func main() {
//
//}
