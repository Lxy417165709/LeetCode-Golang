package 贪心

const INF = 1000000000000

func kLengthApart(nums []int, k int) bool {
	lastPositionOfOne := -INF
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if i-lastPositionOfOne <= k {
			return false
		}
		lastPositionOfOne = i
	}
	return true
}

/*
	题目链接: https://leetcode-cn.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
	总结:
		1. 贪心策略: 遇到1时，只需和上次最近的1的索引做对比。
*/
