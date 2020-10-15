package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := &ListNode{}
	for i := 1; i < 6; i++ {
		fmt.Printf(" insert val:%d\n", i)
		l = addNode(l, i)
	}

	l = removeNthFromEnd(l, 2)
	l = l.Next
	for l != nil {
		fmt.Printf("val:%d\n", l.Val)
		l = l.Next
	}
}

func addNode(l *ListNode, val int) *ListNode {
	h := l
	node := &ListNode{
		Val: val,
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node
	return h
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	p3 := head
	for n > 0 {
		p2 = p2.Next
		n--
	}
	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	del := p1.Next
	p1.Next = del.Next
	return p3
}
