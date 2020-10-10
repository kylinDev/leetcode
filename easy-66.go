package main

import (
	"fmt"
)

func main() {
	// https://github.com/golang/go/wiki/SliceTricks#push-frontunshift
	//[6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,5,4,3] input
	// [6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,4,0,8] error
	// [6,1,4,5,3,9,0,1,9,5,1,8,6,7,0,5,5,4,4] right

	// [7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6]
	// [3,5,4,1,0,2,9,1,0,1,5,6,3,9,8,1,9,9]
	// [7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,7]

	arr := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 9}
	arr = []int{9,9,9}
	res := plusOne(arr)
	fmt.Printf("res:%+v\n", res)
}

func plusOne(digits []int) []int {
	result := make([]int, 0)
	index := len(digits) - 1
	isPlus := true
	var val int
	for index >= 0 {
		if isPlus {
			val = digits[index] + 1
		} else {
			val = digits[index]
		}
		if val > 9 {
			result = append([]int{0}, result...)
		} else {
			result = append([]int{val}, result...)
			isPlus = false
		}
		if index == 0 && isPlus {
			result = append([]int{1}, result...)
			break
		}
		index--
	}
	return result
}
