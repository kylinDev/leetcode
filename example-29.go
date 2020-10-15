package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("min:%0.f,max:%0.f\n", math.Pow(-2, 31), math.Pow(2, 31)-1)
	dividend := -2147483648
	divisor := -1
	result := divide(dividend, divisor)
	fmt.Printf("result:%d\n", result)
}

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	isLessThanZero := false
	if divisor < 0 && dividend < 0 {
		isLessThanZero = false
		dividend = int(math.Abs(float64(dividend)))
		divisor = int(math.Abs(float64(divisor)))
	} else {
		if dividend < 0 {
			isLessThanZero = true
			dividend = int(math.Abs(float64(dividend)))
		}
		if divisor < 0 {
			isLessThanZero = true
			divisor = int(math.Abs(float64(divisor)))
		}
	}
	i := 0
	for dividend > 0 {
		dividend = dividend - divisor
		if dividend >= 0 {
			i++
		}
	}
	if isLessThanZero {
		i = 0 - i
	}
	if i < int(math.Pow(-2, 31)) || i > int(math.Pow(2, 31))-1 {
		return int(math.Pow(2, 31)) - 1
	}
	return i
}
