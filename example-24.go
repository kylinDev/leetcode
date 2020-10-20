package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


func addNode(val int, l *ListNode) *ListNode {
	l1 := l
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node
	return l1
}

func display(l *ListNode) {
	l = l.Next
	for l != nil {
		fmt.Printf("val:%d\n", l.Val)
		l = l.Next
	}
}

func main() {
	l := &ListNode{}
	for i := 0; i < 5; i++ {
		l = addNode(i, l)
	}
	display(l)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	l := head.Next

	return nil
}
