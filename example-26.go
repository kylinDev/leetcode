package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 0, 0, 0, 1, 2, 3, 4, 4, 5, 6, 6, 6, 6, 7, 10, 10}
	size := removeDuplicates(nums)
	fmt.Printf("size:%d,after :%+v\n", size, nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] == nums[i] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	total := len(nums[:j+1])
	return total
}
