package main
import (
	"fmt"
)
func main() {
	str := "hello"
	subStr := "pi"
	fmt.Printf("pos is :%d\n",indexPos(str,subStr))
}

func indexPos(haystack string, needle string) int{
	if haystack == needle {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}
	offset := len(needle)
	pos := len(haystack) - len(needle) +1
	for i:=0;i<pos;i++ {
		target := haystack[i:i+offset]
		fmt.Printf("target:%s\n",target)
		if target == needle {
			return i
		}
	}
	return -1
}