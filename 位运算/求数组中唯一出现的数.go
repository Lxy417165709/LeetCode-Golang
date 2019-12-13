package main

/*
	题目描述:
		给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}
	return ans

	// 这用到了异或运算的性质:
	//					(1) a ^ a == 0
	//					(2) a ^ 0 == a
	//					(3) 异或运算满足结合律、交换律
}

/*
	题目链接: https://leetcode-cn.com/problems/single-number/
*/
