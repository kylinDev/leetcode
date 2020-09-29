package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world  dsf      sdf sfsd  "
	len := lengthOfLastWord(s)
	fmt.Printf("len:%d\n", len)
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	arr := strings.Split(s, " ")
	size := len(arr)
	for size > 0 {
		w := strings.TrimSpace(arr[size-1])
		wordSize := len(w)
		if wordSize > 0 {
			return wordSize
		}
		size--
	}
	return 0
}
