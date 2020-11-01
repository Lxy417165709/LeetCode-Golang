package 位运算

func singleNumber(nums []int) int {
	number := 0
	for i:=0;i<len(nums);i++{
		number^=nums[i]
	}
	return number
}

/*
	题目链接: https://leetcode-cn.com/problems/single-number/
*/
