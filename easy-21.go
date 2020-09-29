/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("main")
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := &ListNode{
	}
	for _, d := range data {
		l.addNode(d)
	}
	l.display()

	data2 := []int{10, 11, 12, 13, 14}
	l2 := &ListNode{
	}
	for _, d := range data2 {
		l2.addNode(d)
	}
	l2.display()
}

func (l *ListNode) addNode(val int) {
	n := l
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	for n.Next != nil {
		n = n.Next
	}
	n.Next = node
}

func (l *ListNode) display() {
	data := l.Next
	for data != nil {
		fmt.Printf("val:%d\n", data.Val)
		data = data.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 := l1.Next
	l2 := l2.Next
	for data1 := l1.Next {
		for data2 := l2.Next {
			if data1.Val <= data2.Val {
				//
				break
			}
		}
	}
	return nil
}
