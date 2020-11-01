package 动态规划


func massage(nums []int) int {
	nums = getExtendArray(nums,2,0)
	maxReservationTime := make([]int,len(nums))
	maxReservationTime[0],maxReservationTime[1]=nums[0],max(nums[0],nums[1])
	for i:=2;i<len(nums);i++{
		maxReservationTime[i] = max(maxReservationTime[i-1],maxReservationTime[i-2]+nums[i])
	}
	return maxReservationTime[len(nums)-1]
}
func getExtendArray(nums []int,lengthOfExtendTo int,fillNum int) []int{
	originLength := len(nums)
	for i:=0;i<lengthOfExtendTo-originLength;i++{
		nums = append(nums,fillNum)
	}
	return nums
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

/*
	题目链接: https://leetcode-cn.com/problems/the-masseuse-lcci/comments/
	总结
		1. 这题和打家劫舍的题是一样的。
			https://leetcode-cn.com/problems/house-robber/
*/
