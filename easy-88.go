package main

import (
	"fmt"
)

func main() {
	fmt.Println("main")
	//nums1 = [1,2,3,0,0,0], m = 3
	//nums2 = [2,5,6],       n = 3
	//merge(nums1, 3, nums2, 3)
	nums1 := []int{0, 0}
	nums2 := []int{0, 0}
	nums := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("nums:%+v\n", nums)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		target := nums2[i]
		isLast := true
		for j := 0; j < m; j++ {
			if target < nums1[j] {
				k := m
				for k > j {
					nums1[k] = nums1[k-1]
					k--
				}
				nums1[j] = target
				m++
				isLast = false
				break
			}
		}
		if isLast {
			nums1[m] = target
			m++
		}
	}
}

