package 二叉树

var countOfSum map[int]int

func findFrequentTreeSum(root *TreeNode) []int {
	countOfSum = make(map[int]int)
	getSumAndFormCountOfSum(root)
	result := make([]int, 0)
	maxCount := 0
	for sum, count := range countOfSum {
		switch {
		case count == maxCount:
			result = append(result, sum)
		case count > maxCount:
			maxCount = count
			result = []int{sum}
		}
	}
	return result
}

func getSumAndFormCountOfSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val + getSumAndFormCountOfSum(root.Left) + getSumAndFormCountOfSum(root.Right)
	countOfSum[sum]++
	return sum
}

/*
	题目链接: https://leetcode-cn.com/problems/most-frequent-subtree-sum/solution/
*/
