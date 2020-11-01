package 前缀和

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		prefixSum: getPrefixSum(nums),
	}
}

func (na *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return na.prefixSum[j]
	}
	return na.prefixSum[j] - na.prefixSum[i-1]
}

func getPrefixSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	return prefixSum

}

/*
	题目链接: https://leetcode-cn.com/problems/range-sum-query-immutable/submissions/
	总结:
		1. 如果数组可修改，那可以最好采用线段树AC这题。
*/
