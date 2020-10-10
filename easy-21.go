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
	data := []int{1,8, 11}
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
	mergeTwoLists(l, l2)
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
	link1 := l1.Next
	link2 := l2.Next
	for link1 != nil {
		isLastOne := true
		for link2 != nil {
			//fmt.Printf("link2 val:%d\n",link2.Val)
			if link1.Val < link2.Val {
				node := &ListNode{
					Val:  link1.Val,
					Next: nil,
				}
				node.Next = link2.Next
				link2 = node
				isLastOne = false
				break
			}
			link2 = link2.Next
		}
		if isLastOne {
			link2 = link1
		}
		link1 = link1.Next
	}
	return link2
}


