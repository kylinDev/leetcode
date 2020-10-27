package main

import (
	"fmt"
)

//输入：s = 7, nums = [2,3,2,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。

func main() {
	nums := []int{2,3,1,2,4,3}
	s := 7
	min := minSubArrayLen(s, nums)
	fmt.Printf("min:%d\n", min)
}

func minSubArrayLen(s int, nums []int) int {
	min := 0
	setMin := false
	for i := 0; i < len(nums); i++ {
		j := i
		sum := 0
		for j < len(nums) {
			sum = sum + nums[j]
			if sum >= s {
				offset := j - i + 1
				if !setMin {
					min = offset
					setMin = true
				}else{
					if offset < min {
						min = offset
					}
				}
				break
			}
			j++
		}
	}
	return min
}
