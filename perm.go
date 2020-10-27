package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	// pass array position start , position end
	perm(nums, 0, 2)
}

func swap(list []int, a int, b int) {
	tmp := list[a]
	list[a] = list[b]
	list[b] = tmp
}

func perm(list []int, k int, m int) {
	//list 数组存放排列的数，K表示层 代表第几个数，m表示数组的长度
	if (k == m) {
		//K==m 表示到达最后一个数，不能再交换，最终的排列的数需要输出；
		for i := 0; i <= m; i++ {
			fmt.Printf("%d, ", list[i])
		}
		fmt.Println("")
	} else {
		for i := k; i <= m; i++ {
			fmt.Printf("swap i=%d,k=%d,list=%+v\n",i,k,list)
			swap(list, i, k)
			fmt.Printf("perm i=%d,k=%d,list=%+v\n",i,k+1,list)
			perm(list, k+1, m)
			fmt.Printf("swap i=%d,k=%d,list=%+v\n",i,k,list)
			swap(list, i, k)
		}
	}
}
