package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := initList()
	addNode(1, l)
	//addNode(2, l)
	//addNode(3, l)
	//addNode(4, l)
	//addNode(5, l)
	//showList(l)
	l = rotateRight(l, 1)
	//showList(l)

	//ll := &ListNode{}
	var ll *ListNode
	ll = addNodeWithoutNilNode(1, ll)
	ll = addNodeWithoutNilNode(2, ll)
	ll = addNodeWithoutNilNode(3, ll)
	ll = addNodeWithoutNilNode(4, ll)
	ll = addNodeWithoutNilNode(5, ll)
	//ll = rotateRight2(ll, 2)
	showList2(ll)
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	var p, prev *ListNode
	for k > 0 {
		p = head
		for head.Next != nil {
			prev = head
			head = head.Next
		}
		prev.Next = nil
		head.Next = p.Next
		// 创建一个头结点
		h := &ListNode{}
		h.Next = head
		head = h
		k--
	}
	return head
}
func rotateRight2(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	q := head
	for q.Next == nil {
		return head
	}
	i := 1
	for q.Next != nil {
		q = q.Next
		i++
	}
	k = k % i
	var h, prev *ListNode
	for k > 0 {
		h = head
		if head.Next == nil {
			return head
		}
		for head.Next != nil {
			prev = head
			head = head.Next
		}
		prev.Next = nil
		head.Next = h
		k--
	}
	return head
}

// 无头结点
func addNode2(val int, l *ListNode) *ListNode {
	if l == nil {
		l = &ListNode{
			Val: val,
		}
		return l
	}
	// 保持第一个元素地址
	head := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{
		Val: val,
	}
	return head
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


func showList2(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func showList(l *ListNode) {
	// 排除头结点
	l = l.Next
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func initList() *ListNode {
	return &ListNode{
	}
}

func addNode(val int, l *ListNode) *ListNode {
	for l.Next != nil {
		l = l.Next
	}
	l.Next = &ListNode{
		Val: val,
	}
	return l
}
