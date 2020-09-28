package main
import (
	"fmt"
)
func main(){
	nums :=[]int{1,3,5,6}
	target := 0
	pos := searchInsert(nums,target)
	fmt.Printf("pos:%d\n",pos)
}

func searchInsert(nums []int, target int) int {

	for i:=0;i<len(nums);i++ {
		if nums[i]>= target {
			return i
		}
	}
	return len(nums)
}