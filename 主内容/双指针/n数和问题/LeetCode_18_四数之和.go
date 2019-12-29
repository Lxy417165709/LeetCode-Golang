package 双指针

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	// mA,mB是为了去重
	sort.Ints(nums)
	fmt.Println(nums)
	mA:=make(map[int]int)
	slice:=make([][]int,0)
	for i:=0;i< len(nums);i++  {
		mB:=make(map[int]int)	// 必须放在里面，因为被使用时相对于i的
		if mA[nums[i]]==1 {
			continue
		}
		mA[nums[i]] = 1
		for t:=i+1;t<len(nums) ;t++  {
			if mB[nums[t]]==1 {
				continue
			}
			mB[nums[t]] = 1
			sum:=target - nums[i]-nums[t]
			l,r:=t+1, len(nums)-1

			for ;l<r ;  {
				tmp:=nums[l]+nums[r]
				if sum==tmp {
					slice = append(slice, []int{
						nums[i],nums[t],nums[l],nums[r],
					})
					a,b:=nums[l],nums[r]
					for ;l<r && nums[l]==a ;  {
						l++
					}
					for ;l<r && nums[r]==b ;  {
						r--
					}
				}else{
					if sum>tmp {
						l++
					}else{
						r--
					}
				}
			}
		}
	}
	return slice
}

func main() {
	fmt.Println(fourSum([]int{
		0,0,0,0,
	},0))
}
/*
	总结
	1. 我解答该题时间复杂度O(n^3)，先选2个数(注意上面的去重代码)，之后两数和使用O(n)算法
	2. 官方有时间复杂度O(n^2),思想是使用2次twoSum
*/