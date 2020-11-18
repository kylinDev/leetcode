package main

import (
	"fmt"
	"strings"
)

func main() {

	// egg add
	//"abab"
	//"baba"
	s := "abab"
	t := "baba"
	result := isIsomorphic(s, t)
	fmt.Printf("result:%+v\n", result)
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[byte]byte,0)
	for i := 0; i < len(t); i++ {
		if _,ok := m[t[i]];ok {
			continue
		}
		m[t[i]] = t[i]
		s = strings.Replace(s, string(s[i]), string(t[i]), -1)
	}
	return s == t
}
