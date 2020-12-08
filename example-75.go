package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 1, 2, 0, 0, 2, 1, 1, 0, 1}
	sortColors(nums)
}

func sortColors(nums []int) {
	j := len(nums) - 1
	i := 0
	for i != j {
		fmt.Printf("i:%d,j:%d\n", i, j)
		if nums[i] > nums[j] {
			v := nums[j]
			nums[j] = nums[i]
			nums[i] = v
			j--
			i++
		} else if nums[i] == nums[j] {
			j--
		} else {
			i++
		}
	}
	fmt.Printf("nums:%+v\n", nums)
}
