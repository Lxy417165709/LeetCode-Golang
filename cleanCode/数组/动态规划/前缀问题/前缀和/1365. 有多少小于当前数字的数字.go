package 前缀和

// ------------------------- 哈希 + 前缀和 -------------------------
// 执行用时：4 ms,   在所有 Go 提交中击败了 94.44% 的用户
// 内存消耗：3.3 MB, 在所有 Go 提交中击败了 100.00% 的用户
func smallerNumbersThanCurrent(nums []int) []int {
	countOfNum := getCountOfNum(nums)
	countOfLessOrEqualNum := getPrefixSum(countOfNum)
	return getArrayOfCountOfLessNum(nums, countOfLessOrEqualNum)
}

func getArrayOfCountOfLessNum(nums []int, countOfLessOrEqualNum []int) []int {
	// 其实可以把结果存在 nums 中，这样能节省一点空间，但也会带来一个问题: nums被修改了。
	array := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 >= 0 {
			array[i] = countOfLessOrEqualNum[nums[i]-1]
		} else {
			array[i] = 0
		}
	}
	return array
}

func getPrefixSum(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	prefixSum := make([]int, len(arr))
	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}

func getCountOfNum(nums []int) []int {
	countOfNum := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		countOfNum[nums[i]]++
	}
	return countOfNum
}
