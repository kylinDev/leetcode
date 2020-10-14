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
	l := &ListNode{
	}
	for _, d := range data {
		l.addNode(d)
	}
	//l.display()

	data2 := []int{6, 9, 10, 13, 14}
	l2 := &ListNode{
	}
	for _, d := range data2 {
		l2.addNode(d)
	}
	//l2.display()
	mergeLink := mergeTwoLists(l, l2)

	for mergeLink != nil {
		fmt.Printf("mergeLink val:%d\n", mergeLink.Val)
		mergeLink = mergeLink.Next
	}
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
	link1 := l1
	link2 := l2
	link3 := l2
	for link1.Next != nil {
		for link2.Next != nil {
			if link1.Val < link2.Val {
				fmt.Printf("insert %d before %d\n", link1.Val, link2.Val)
				node := &ListNode{
					Val:  link1.Val,
					Next: nil,
				}
				node.Next = link2
				link2 = node
				break
			}
			link2 = link2.Next
		}
		fmt.Printf("l1 val:%d\n", link1.Val)
		link1 = link1.Next
	}
	return link3
}
