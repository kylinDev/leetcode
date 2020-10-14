package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("v:%0.f,v2:%0.f,s:%s\n", math.Pow(-2, 31), math.Pow(2, 31)-1,strconv.Itoa(-120))
	v := 1534236469
	rs := reverse(v)
	fmt.Printf("rs:%d\n", rs)
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	if isOverflow(x) {
		return 0
	}
	if x > 0 {
		numsStr := reverseInt(x)
		result, err := strconv.Atoi(numsStr)
		if err != nil {
			return 0
		}
		if isOverflow(result) {
			return 0
		}
		return result
	} else {
		x = 0 - x
		numsStr := reverseInt(x)
		result, err := strconv.Atoi(numsStr)
		if err != nil {
			return 0
		}
		result = 0 - result
		if isOverflow(result) {
			return 0
		}
		return result
	}
}

func isOverflow(x int) bool {
	if float64(x) > math.Pow(2, 31)-1 || float64(x) < math.Pow(-2, 31) {
		return true
	}
	return false
}

func reverseInt(x int) string {
	nums := strings.Split(strconv.Itoa(x), "")
	size := len(nums) - 1
	result := make([]string, 0)
	for size >= 0 {
		if size == 0 && nums[size] == "0" {
			break
		}
		result = append(result, nums[size])
		size--
	}
	return strings.Join(result, "")
}
