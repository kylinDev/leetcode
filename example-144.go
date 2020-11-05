package main
import (
	"fmt"
	_"sync"
)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	 Root *TreeNode
}

func main() {
	t := &Tree{
	}
	nums := []int{1,2,3,68,0,8}
	for _, num := range nums {
		t.insert(num)
	}
	t.beginSort(t.Root)

	result := preorderTraversal(t.Root)
	fmt.Printf("result:%+v\n",result)
}

//二叉搜索树,左值小于右值
func (t *Tree) insert(val int) {
	node := &TreeNode{
		Val: val,
	}
	if t.Root == nil {
		t.Root = node
		return
	}
	cur := t.Root
	for {
		if cur.Val > node.Val {
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

func (t *Tree) beginSort(node *TreeNode) {
	if node != nil {
		fmt.Printf("beginSort val:%d\n", node.Val)
		t.beginSort(node.Left)
		t.beginSort(node.Right)
	}
}

// 递归
func preorderTraversal(root *TreeNode) []int {
	result := make([]int,0)
	type traversal func(*TreeNode)
	var f traversal
	f = func(root *TreeNode) {
		if root != nil {
			result = append(result,root.Val)
			f(root.Left)
			f(root.Right)
		}
	}
	f(root)
	return result
}



// 非递归
func preOrderNR(node *Node) []int {
	var result []int
	stack := make([]*Node, 0)
	stack = append(stack, node)
	for len(stack) > 0 {
		size := len(stack)
		// 2句为出栈操作
		curNode := stack[size-1]
		stack = stack[0 : size-1]
		result = append(result, curNode.Value)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	return result
}

