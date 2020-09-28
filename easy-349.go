package main

import "fmt"

func main() {
	//nums1 = [1,2,2,1], nums2 = [2,2]
	// nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}

	result := intersection(nums1, nums2)
	fmt.Printf("result:%+v\n", result)
}

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	uniMap := make(map[int]int,0)
	for i := 0; i < len(nums1); i++ {
		isIn := inArr( nums1[i], nums2)
		if isIn {
			if _,ok:=uniMap[nums1[i]];!ok {
				result = append(result, nums1[i])
				uniMap[nums1[i]] = nums1[i]
			}
		}
	}
	return result
}

func inArr(v int, target []int) bool {
	for i := 0; i < len(target); i++ {
		if v == target[i] {
			return true
		}
	}
	return  false
}
