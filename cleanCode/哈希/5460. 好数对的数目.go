package 哈希

func numIdenticalPairs(nums []int) int {
	preCount := make(map[int]int)
	countOfGoodPair := 0
	for i := 0; i < len(nums); i++ {
		countOfGoodPair += preCount[nums[i]]
		preCount[nums[i]]++
	}
	return countOfGoodPair
}

/*
	题目链接: https://leetcode-cn.com/problems/number-of-good-pairs/
*/
