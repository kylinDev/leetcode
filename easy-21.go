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
	//fmt.Println("main")


	data := []int{1, 7, 19}
	var l *ListNode
	for _, d := range data {
		l = addNode(d, l)
	}
	//display(l)

	data2 := []int{2, 3, 4}
	var l2 *ListNode
	for _, d := range data2 {
		l2 = addNode(d, l2)
	}
	//display(l2)

	l3 := mergeTwoLists(l,l2)

	display(l3)
}

func addNode(val int, l *ListNode) *ListNode {
	node := &ListNode{
		Val: val,
	}
	if l == nil {
		l = node
		return l
	}
	var head *ListNode
	head = l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node
	return head
}

func display(l *ListNode) {
	for l != nil {
		fmt.Printf("val:%d\n", l.Val)
		l = l.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val <= l2.Val {
		head = l1
	}else{
		head = l2
	}

	for l1 != nil {
		for l2 != nil {
			n := &ListNode{
				Val: l1.Val,
			}
			if l1.Val >= l2.Val {
				l2.Next = n
				n.Next = l2
			}else{
				n.Next = l2
			}
			l1 = l1.Next
			break
		}
	}

	return head
}
