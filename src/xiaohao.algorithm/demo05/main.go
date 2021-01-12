package main

import(
	"fmt"
)

//27.移除元素
func removeElement(nums []int,val int)int{
	for i:=0;i<len(nums);{
	if nums[i]==val{
			nums = append(nums[:i],nums[i+1:]...)
		}else{
			i++
		}
	}
	return len(nums)
}

func main(){
	nums:=[]int{3,2,2,3}
	val:=3
	fmt.Println(removeElement(nums,val))
}