package 模拟

func decompressRLElist(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		result = insert(result, nums[i], nums[i+1])
	}
	return result
}

func insert(array []int, times int, num int) [] int {
	for i := 0; i < times; i++ {
		array = append(array, num)
	}
	return array
}

/*
	题目链接: https://leetcode-cn.com/problems/decompress-run-length-encoded-list/submissions/
*/
