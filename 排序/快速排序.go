package main



// 快速排序
func quickSort(nums []int){
	if len(nums)==0{
		return
	}
	l,r:=0,len(nums)-1
	for l<=r{
		for l<=r && nums[l]<=nums[0]{
			l++
		}
		for l<=r && nums[r]>=nums[0]{
			r--
		}
		if l<=r {
			nums[l],nums[r] = nums[r],nums[l]
		}
	}
	nums[0],nums[r] = nums[r],nums[0]

	quickSort(nums[:r])
	quickSort(nums[r+1:])
}




/*
	题目链接: https://leetcode-cn.com/problems/sort-an-array/
*/