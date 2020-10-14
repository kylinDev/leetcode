package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	_ "unicode"
)

func main() {
	s := "9223372036854775808"
	i, e := strconv.Atoi(s)
	fmt.Printf("i:%+v,e:%+v\n", i, e)
	fmt.Printf("v:%s,max pow:%0.f,min pow:%0.f\n", s, math.Pow(2, 31)-1, math.Pow(-2, 31))
	nums := myAtoi(s)
	fmt.Printf("nums:%d\n", nums)
}

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	result := make([]string, 0)
	for i, ch := range []byte(s) {
		if string(ch) == "+" || string(ch) == "-" {
			if i >= 1 {
				return 0
			}
			continue
		}
		if ch < '0' || ch > '9' {
			break
		}
		result = append(result, string(ch))
	}
	if len(result) == 0 {
		return 0
	}
	isMinusSign := false
	if s[0] == []byte("-")[0] {
		isMinusSign = true
	}
	nums := toNums(strings.Join(result, ""), isMinusSign)
	return nums
}

func toNums(s string, isMinusSign bool) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)
	}
	if isMinusSign {
		n = 0 - n
		if n < int(math.Pow(-2, 31)) {
			return int(math.Pow(-2, 31))
		}
	}
	if n > int(math.Pow(2, 31))-1 {
		return int(math.Pow(2, 31)) - 1
	}
	return n
}
