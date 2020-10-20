package main

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	findMedianSortedArrays(nums1, num2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	for i := 1; i < len(nums1); i++ {
		target := nums1[i]
		j := i
		for j > 0 && nums1[j-1] > target {
			nums1[j] = nums1[j-1]
			j--
		}
		nums1[j] = target
	}
	size := len(nums1)
	if size == 1 {
		return float64(nums1[0])
	}
	if size%2 == 0 {
		pre := size/2 - 1
		return float64((nums1[pre] + nums1[pre+1])) / float64(2)
	}
	return float64(nums1[size/2])
}
