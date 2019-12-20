package main

/*
	给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。
*/
func subarraySum(nums []int, k int) int {
	sum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	ans := 0
	count := make(map[int]int)
	for i := 0; i < len(sum); i++ {
		ans += count[sum[i]-k]
		count[sum[i]]++
	}
	return ans
}

/*
	题目链接:
		https://leetcode-cn.com/problems/subarray-sum-equals-k/		和为K的子数组
*/

/*
	总结
	1. 这题目可以使用暴力，枚举2个端点，之后计算总和，但是时间复杂度为O(n^3)，其中枚举O(n^2)，计算总和O(n)
	2. 上面计算总和用了O(n)，于是可以使用前缀和数组记录总和，将计算总和优化为O(1),
		于是可以把暴力法时间复杂度优化为O(n^2)。
	3. 第三种方法的时间复杂度是O(n)，思路就是先记录前缀和，之后再从左到右扫描前缀和数组，记录到目前为止该前缀和
		出现了多少次(计为count[sum[i]])，那么count[sum[i]-k]就是以nums[i]结尾且和为k的子数组的个数，累加起来就可以了。
*/
