package main

import (
	"fmt"
	"math"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

type Tree struct {
	Root *Node
}

func main() {
	t := &Tree{
	}
	nums := []int{15, 6, 23, 7, 4, 71, 5, 50}
	for _, num := range nums {
		t.insert(num)
	}
	//t.beginSort(t.Root)
	t.middleSort(t.Root)
	t.endSort(t.Root)
	//fmt.Printf("find %d in tree ?,%+v\n", 2, t.find(2))
	preOrderNR(t.Root)
	MidOrderNR(t.Root)
	t.endSort(t.Root)
	endOrderRN(t.Root)
	//max := maxDepth(t.Root)
	//fmt.Printf("max:%d\n", max)
}

func (t *Tree) find(target int) bool {
	cur := t.Root
	if cur == nil {
		return false
	}
	for cur.Value != target {
		if cur.Value < target {
			cur = cur.Right
		} else {
			cur = cur.Left
		}
		if cur == nil {
			return false
		}
	}
	return true
}

//二叉搜索树,左值小于右值
func (t *Tree) insert(val int) {
	node := &Node{
		Value: val,
	}
	if t.Root == nil {
		t.Root = node
		return
	}
	cur := t.Root
	for {
		if cur.Value > node.Value {
			if cur.Left == nil {
				cur.Left = node
				return
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = node
				return
			}
			cur = cur.Right
		}
	}
}

func (t *Tree) beginSort(node *Node) {
	if node != nil {
		fmt.Printf("beginSort val:%d\n", node.Value)
		t.beginSort(node.Left)
		t.beginSort(node.Right)
	}
}

func (t *Tree) middleSort(node *Node) {
	if node != nil {
		t.middleSort(node.Left)
		fmt.Printf("middleSort val:%d\n", node.Value)
		t.middleSort(node.Right)
	}
}

func (t *Tree) endSort(node *Node) {
	if node != nil {
		t.endSort(node.Left)
		t.endSort(node.Right)
		fmt.Printf("endSort val:%d\n", node.Value)
	}
}

func preOrderNR(node *Node) []int {
	var result []int
	stack := make([]*Node, 0)
	// 根入栈
	stack = append(stack, node)
	for len(stack) > 0 {
		size := len(stack)
		// 2句为出栈操作
		curNode := stack[size-1]
		stack = stack[0 : size-1]
		result = append(result, curNode.Value)
		// 栈是先进后出，所以右节点先入栈
		// 前序:根左右
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	fmt.Printf("result:%+v\n", result)
	return result
}

func MidOrderNR(node *Node) {
	var result []int
	var stack []*Node
	if node == nil {
		return
	}
	for node != nil {
		stack = append(stack, node)
		node = node.Left
	}
	for len(stack) > 0 {
		size := len(stack)
		node := stack[size-1]     // last
		stack = stack[0 : size-1] // del last
		result = append(result, node.Value)
		if node.Right != nil {
			node = node.Right
			stack = append(stack, node)
			for node.Left != nil {
				node = node.Left
				stack = append(stack, node)
			}
		}
	}
	fmt.Printf("MidOrderNR result:%+v\n", result)
}

func endOrderRN(node *Node) []int {
	var result []int
	var stackOne []*Node
	var stackTwo []*Node
	if node == nil {
		return result
	}
	stackOne = append(stackOne, node)
	for len(stackOne) > 0 {
		size := len(stackOne)
		n := stackOne[size-1]
		stackOne = stackOne[0 : size-1]
		stackTwo = append(stackTwo, n)
		if n.Left != nil {
			stackOne = append(stackOne, n.Left)
		}
		if n.Right != nil {
			stackOne = append(stackOne, n.Right)
		}
	}
	for len(stackTwo) > 0 {
		size := len(stackTwo)
		n := stackTwo[size-1]
		stackTwo = stackTwo[0 : size-1]
		result = append(result, n.Value)
	}
	return result
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	if root != nil {
		if !isValidBST(root.Left) {
			return false
		}
		if prev != nil && prev.Val > root.Val {
			return false
		}
		prev = root
		if !isValidBST(prev.Right) {
			return false
		}
	}
	return true
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)
	fmt.Printf("lh:%d,rh:%d\n", lh, rh)
	return int(math.Max(float64(lh), float64(rh))) + 1
}
