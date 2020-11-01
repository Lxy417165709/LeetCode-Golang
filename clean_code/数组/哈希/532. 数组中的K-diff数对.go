package 哈希

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	countOfNum := getCountOfNum(nums)
	result := 0
	for num, _ := range countOfNum {
		if num-k == num && countOfNum[num-k] >= 2 {
			result++
		}
		if num-k != num && countOfNum[num-k] >= 1 {
			result++
		}
	}
	return result
}

func getCountOfNum(nums []int) map[int]int {
	countOfNum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countOfNum[nums[i]]++
	}
	return countOfNum
}


/*
	题目链接: https://leetcode-cn.com/problems/k-diff-pairs-in-an-array/submissions/
*/
