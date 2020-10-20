package main

import (
	"fmt"
)

func main() {
	s := "pwwkew"
	max := lengthOfLongestSubstring(s)
	fmt.Printf("max:%d\n", max)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		j := i
		size := 0
		counter := make(map[string]string)
		for j < len(s) {
			w := string(s[j])
			if _, ok := counter[w]; ok {
				break
			} else {
				counter[w] = w
			}
			size++
			j++
		}
		if size > max {
			max = size
		}
	}
	return max
}
