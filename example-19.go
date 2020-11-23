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
	Prev *ListNode
}

func main() {
	//l := &ListNode{}
	//for i := 1; i < 6; i++ {
	//	//fmt.Printf(" insert val:%d\n", i)
	//	l = addNode(l, i)
	//}
	//
	//l = removeNthFromEnd(l, 2)
	//l = l.Next
	//for l != nil {
	//	//fmt.Printf("val:%d\n", l.Val)
	//	l = l.Next
	//}

	var ll *ListNode

	ll = addNodeWithoutNilNode(1, ll)
	ll = addNodeWithoutNilNode(2, ll)
	//ll = addNodeWithoutNilNode(3, ll)
	//ll = addNodeWithoutNilNode(4, ll)
	//ll = removeNthFromEnd(ll, 1)
	//displayList(ll)

	var l *ListNode
	l = addReverseNode(6, l)
	l = addReverseNode(7, l)
	l = addReverseNode(8, l)
	l = addReverseNode(9, l)
	//displayList(l)

	var l2 *ListNode

	l2 = addReverseNode(3, l2)
	l2 = addReverseNode(4, l2)
	l2 = addReverseNode(5, l2)
	displayList(l2)
	l2 = reverseList(l2)
	displayList(l2)


	var l3 *ListNode
	l3 = addDouleListNode(6,l3)
	l3 = addDouleListNode(9,l3)
	l3 = addDouleListNode(19,l3)
	l3 = addDouleListNode(129,l3)
	l3 = addDouleListNode(10,l3)
	showDoubleList(l3)
}


func showDoubleList(l *ListNode) {
	for l != nil {
		fmt.Printf("curr node val:%d\n",l.Val)
		if l.Prev != nil {
			fmt.Printf("prev node val:%d\n",(l.Prev).Val)
		}else{
			fmt.Printf("no prev node\n")
		}
		if l.Next != nil {
			fmt.Printf("next node val:%d\n",(l.Next).Val)
		}else{
			fmt.Printf("no next node\n")
		}
		l = l.Next
	}
}




func addDouleListNode(val int,l *ListNode) *ListNode {
	n := &ListNode{
		Val:val,
	}
	if l == nil {
		l = n
		return l
	}
	var head *ListNode
	head = l
	for l.Next != nil {
		l = l.Next
	}
	n.Prev = l
	l.Next = n
	return head
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

func displayList(l *ListNode) {
	for l != nil {
		fmt.Printf("%d\n", l.Val)
		l = l.Next
	}
}

//反转单链表
func reverseList(l *ListNode) *ListNode {
	var ll *ListNode
	for l != nil {
		p := &ListNode{
			Val: l.Val,
		}
		if ll == nil {
			ll = p
		} else {
			p.Next = ll
			ll = p
		}
		l = l.Next
	}
	return ll
}

// 逆序增加节点
func addReverseNode(val int, l *ListNode) *ListNode {
	if l == nil {
		l = &ListNode{
			Val: val,
		}
		return l
	}
	p := &ListNode{
		Val: val,
	}
	p.Next = l
	return p
}

// 无空节点链表
func addNodeWithoutNilNode(val int, l *ListNode) *ListNode {
	if l == nil {
		l = &ListNode{
			Val: val,
		}
		return l
	}
	head := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{
		Val: val,
	}
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	p2 := head
	p3 := head
	if head == nil {
		return head
	}
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
