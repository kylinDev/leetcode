package main

import (
	"fmt"
	"sync"
)

type MinStack struct {
	container []int
	min       int
	lock      sync.Mutex
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.container = append(this.container, x)
	this.setMin()
}

func (this *MinStack) Pop() {
	if len(this.container) == 0 {
		return
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	c := make([]int, 0)
	size := len(this.container)
	for i := 0; i <= size-2; i++ {
		c = append(c, this.container[i])
	}
	this.container = c
	if this.container[len(this.container)] == this.min {
		this.setMin()
	}
}

func (this *MinStack) setMin() {
	if len(this.container) == 0 {
		return
	}
	min := this.container[0]
	for i := 1; i < len(this.container); i++ {
		if min > this.container[i] {
			min = this.container[i]
		}
	}
	this.min = min
}

func (this *MinStack) Top() int {
	size := len(this.container)
	if size == 0 {
		return -1
	}
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.container[size-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	//minStack.push(-2);
	//minStack.push(0);
	//minStack.push(-3);
	//minStack.getMin();   --> 返回 -3.
	//minStack.pop();
	//minStack.top();      --> 返回 0.
	//minStack.getMin();   --> 返回 -2.
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	//obj.Push(-5)
	min := obj.GetMin()
	fmt.Printf("min=%d\n", min)
	obj.Pop()
	top := obj.Top()
	fmt.Printf("top=%d\n", top)
	min = obj.GetMin()
	fmt.Printf("min=%d\n", min)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
