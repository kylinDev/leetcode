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
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = l1.Next
	l2 = l2.Next
	l3 := &ListNode{}
	for l1 != nil && l2 != nil {
		pluse := false
		fmt.Printf("v1:%d,v2:%d\n", l1.Val, l2.Val)
		l1 = l1.Next
		l2 = l2.Next
		if (l1.Val + l2.Val) > 9 {
			pluse = true
		} else {
			sum := l1.Val + l2.Val
			node := &ListNode{
				l1.Val + l2.Val,
			}
		}
	}
	return nil
}

func addReverseNode(val int, l *ListNode) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	l = l.Next
	if l == nil {
		l.Next = node
		return l
	}
	node.Next = l
	return node
}

func addNode(val int, l *ListNode) *ListNode {
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

func display(l *ListNode) {
	l = l.Next
	for l != nil {
		fmt.Printf("val:%d\n", l.Val)
		l = l.Next
	}
}

func main() {
	l := &ListNode{
	}
	ll := &ListNode{
	}
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}
	for i := 0; i < len(nums1); i++ {
		fmt.Printf("i:%d\n", nums1[i])
		l = addNode(nums1[i], l)
	}
	//display(l)

	for i := 0; i < len(nums2); i++ {
		fmt.Printf("i:%d\n", nums2[i])
		ll = addNode(nums2[i], ll)
	}
	//display(ll)
	l2 := &ListNode{}
	for i := 1; i < 5; i++ {
		l2 = addReverseNode(i, l2)
	}
	display(l2)
	//addTwoNumbers(l, ll)
}
