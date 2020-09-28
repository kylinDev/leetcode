package main

import (
	"fmt"
	_"strings"
)
func main(){
	testCase :=[]string{
		"{}","[]","()","{]","{[}]","([)]","{[]}",
	}

	for _,v := range testCase{
		is := isBraket(v)
		fmt.Printf("%s is %+v\n",v,is)
	}
}

// () {} []
func isBraket(str string) bool{
	stack := make([]string,0)
	for i:=0;i<len(str);i++ {
		w := string(str[i])
		if w=="(" || w == "{" || w == "[" {
			stack = append(stack,w)
		}
		if (w == ")" || w == "}" || w == "]") && len(stack) == 0{
			return false
		}
		size := len(stack)
		if w == ")" {
			top := stack[size-1]
			if top != "(" {
				return false
			}else{
				stack = stack[:size -1]
			}
		}
		if w == "}" {
			top := stack[size-1]
			if top != "{" {
				return false
			}else{
				stack = stack[:size -1]
			}
		}
		if w == "]" {
			top := stack[size-1]
			if top != "[" {
				return false
			}else{
				stack = stack[:size -1]
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}