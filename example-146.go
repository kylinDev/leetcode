package main
import (
	"fmt"
)
type LRUCache struct {
	capacity int
	m        map[int]int
	l        []int
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]int, capacity)
	return LRUCache{
		capacity: capacity,
		m:        m,
	}
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.m[key]; ok {
		this.findPos(key)
		this.l[len(this.l)-1] = key
		return value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if len(this.m) < this.capacity {
		if _, ok := this.m[key]; !ok {
			this.m[key] = value
			this.l = append(this.l, key)
			return
		} else {
			pos := this.findPos(key)
			delete(this.m, pos)
			this.m[key] = value
			this.l[len(this.l)-1] = key
		}
	} else {
		if _, ok := this.m[key]; !ok {
			delete(this.m, this.l[0])
			this.m[key] = value
			this.l = append(this.l, key)
			this.l = this.l[1:]
			return
		}
		pos := this.findPos(key)
		delete(this.m, pos)
		this.m[key] = value
		this.l[len(this.l)-1] = key
	}
}

// 找到key，逐个往前覆盖
func (this *LRUCache) findPos(key int) int {
	for i := 0; i < len(this.l); i++ {
		j := i + 1
		if this.l[i] == key {
			for j < len(this.l) {
				this.l[i] = this.l[j]
				j++
				i++
			}
			return key
		}
	}

	return -1
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {


	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	// [null,null,null,1,null,-1,null,1,3,4]
	// [null,null,null,1,null,-1,null,-1,3,4]

	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	v := obj.Get(1)
	fmt.Printf("v:%d\n", v)
	obj.Put(3, 3)
	v = obj.Get(2)
	fmt.Printf("v:%d\n", v)
	obj.Put(4, 4)
	v = obj.Get(1)
	fmt.Printf("v:%d\n", v)
	v = obj.Get(3)
	fmt.Printf("v:%d\n", v)
	v = obj.Get(4)
	fmt.Printf("v:%d\n", v)
	//fmt.Printf("l:%+v\n", obj.l)
	//v := obj.Get(2)
	//fmt.Printf("v:%d\n", v)
	//obj.Put(1, 1)
	//obj.Put(4, 1)
	//v = obj.Get(2)
	//fmt.Printf("v:%d\n", v)
}
