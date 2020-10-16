package main

import (
	"fmt"
)

func main() {
	// [3,2,2,3], val = 3,
	//[0,1,2,2,3,0,4,2]
	//2
	nums := []int{3,2,2,3}
	size := removeElement(nums, 3)
	fmt.Printf("nums:%+v,size:%d\n", nums[:size], size)
}

func removeElement(nums []int, val int) int {
	hit := 0
	for i := 0; i < len(nums); i++ {
		j := i
		if nums[i] != val {
			continue
		}else{
			hit++
		}
		for j < len(nums) {
			if nums[j] != val {
				nums[i] = nums[j]
				nums[j] = val
				hit--
				break
			}
			j++
		}
	}
	return len(nums) - hit
}

// bad case,只能删除重复有序的数据
func removeElement2(nums []int, val int) int {
	hit := 0
	for i := 0; i < len(nums); i++ {
		j := i + 1
		if nums[i] == val {
			hit++
		}
		if nums[i] != val {
			continue
		}
		for j < len(nums) {
			if nums[j] != val {
				nums[i] = nums[j]
				break
			}
			j++
		}
		for j < len(nums)-1 {
			nums[j] = nums[j+1]
			j++
		}
	}
	return len(nums) - hit
}
