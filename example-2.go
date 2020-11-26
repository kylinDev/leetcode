package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
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

func pow(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Pow10(n))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum1, sum2 int
	n := 0
	for l1 != nil {
		fmt.Printf("sum2:%d,v:%d\n", pow(n), l1.Val)
		sum1 += pow(n) * l1.Val
		n++
		l1 = l1.Next
	}
	fmt.Printf("sum1:%d\n", sum1)
	m := 0
	for l2 != nil {
		fmt.Printf("sum2:%d,v:%d\n", pow(m), l2.Val)
		sum2 += pow(m) * l2.Val
		m++
		l2 = l2.Next
	}
	fmt.Printf("sum2:%d\n", sum2)
	sum := sum1 + sum2
	s := strings.Split(strconv.Itoa(sum), "")
	var l3 *ListNode
	for i := 0; i < len(s); i++ {
		v, _ := strconv.Atoi(s[i])
		l3 = reverseList(v, l3)
	}
	return l3
}

func reverseList(val int, l *ListNode) *ListNode {
	n := &ListNode{
		Val: val,
	}
	if l == nil {
		l = n
		return l
	}
	n.Next = l
	return n
}

func addReverseNode(val int, l *ListNode) *ListNode {
	node := &ListNode{
		Val:  val,
		Next: nil,
	}
	node.Next = l.Next
	return node
}

func addNode(val int, l *ListNode) *ListNode {
	node := &ListNode{
		Val: val,
	}
	if l == nil {
		l = node
		return l
	}
	h := l
	for l.Next != nil {
		l = l.Next
	}
	l.Next = node
	return h
}

func display(l *ListNode) {
	for l != nil {
		fmt.Printf("val:%d\n", l.Val)
		l = l.Next
	}
}

func main() {
	var l *ListNode
	var l2 *ListNode
	nums1 := []int{2, 4, 3}
	for i := 0; i < len(nums1); i++ {
		l = addNode(nums1[i], l)
	}
	//display(l)

	nums2 := []int{5, 6, 4}
	for i := 0; i < len(nums2); i++ {
		l2 = addNode(nums2[i], l2)
	}
	//display(l2)
	l3 := addTwoNumbers(l, l2)
	display(l3)
	//addTwoNumbers(l, ll)
}
