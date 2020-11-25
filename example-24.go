package main

import (
	"fmt"
	_ "os"
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
	if l == nil {
		l = node
		return l
	}
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node
	return l1
}

func display(l *ListNode) {
	for l != nil {
		fmt.Printf("display val:%d\n", l.Val)
		l = l.Next
	}
}

func main() {
	var l *ListNode
	for i := 1; i < 6; i++ {
		fmt.Printf("insert val:%d\n", i)
		l = addNode(i, l)
	}
	l = swapPairs(l)
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
	if head == nil || head.Next == nil {
		return head
	}
	var q *ListNode
	p := head.Next
	for head.Next != nil {
		next := (head.Next).Next // 第三节点
		if next == nil {
			(head.Next).Next = head
			if q != nil {
				q.Next = head.Next
			}
			head.Next = nil
			break
		}
		(head.Next).Next = head  // 第二节点连接第一节点
		head.Next = next         // 第一节点连接第三节点
		q = head
		head = next // 指针滑动到交换后第三节点
	}
	return p
}
