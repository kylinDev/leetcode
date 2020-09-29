package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "h"
	len := lengthOfLastWord(s)
	fmt.Printf("len:%d\n", len)
}

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	size := len(s)
	isStart := false
	wordSize := 0
	for size > 0 {
		w := string(s[size-1])
		if !unicode.IsSpace([]rune(w)[0]) {
			isStart = true
			wordSize++
		}
		if isStart && unicode.IsSpace([]rune(w)[0])  {
			break
		}
		size--
	}
	return wordSize
}
