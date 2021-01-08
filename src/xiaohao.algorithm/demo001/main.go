package main

import(
	"sort"
	"fmt"
)

//两个数组的交集
func intersect(nums1 [4] int,nums2 [5] int)[]int{
	m0:=map[int]int{}
	for _,v :=range nums1{
		//遍历nums1,初始化map
		m0[v]+=1
	}
	k:=0
	for _,v:=range nums2{
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v]>0{
			m0[v]-=1
			nums2[k]=v
			k++
		}
	}
	return nums2[0:k]
}

//两个数组的交集,排序好的数据
func intersect2(nums1 []int,nums2[]int)[]int{
	i,j,k:=0,0,0
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i<len(nums1) && j<len(nums2){
		if nums1[i]>nums2[j]{
			j++
		}else if nums1[i]<nums2[j]{
			i++
		}else{
			nums1[k]=nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}

func main(){

	nums1:=[4]int{2,3,8,1}
	nums2:=[5]int{8,1,9,0,4}

	res:=intersect(nums1,nums2)
	fmt.Println(res)

}