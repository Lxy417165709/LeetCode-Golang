package 一维数组

import "sort"

// nextPermutation 找到下一个排列。
// 解释: 把 nums 看做数字，比如 [1, 3, 2] 看做 132。
//      下一个排列就是: 第一个大于 132 的数值对应的数组，如 [2, 1, 3]，而 [2, 3, 1] 是第二个。
func nextPermutation(nums []int) {
	// 1. 逆序获取落点索引。
	dropIndex := -1 // 逆序非递减遇到的第一个落点索引。
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			dropIndex = i - 1
			break
		}
	}

	// 2. 没有落点，说明排列已是最大。
	if dropIndex == -1 {
		sort.Ints(nums)
		return
	}

	// 3. 逆序获取第一个大于 落点索引对应数字 的索引。
	indexOfFirstGreaterDropNum := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] > nums[dropIndex] {
			indexOfFirstGreaterDropNum = i
			break
		}
	}

	// 4. 交换、递增排序落点索引后的元素(保证低位数字大、高位数字小)。
	nums[indexOfFirstGreaterDropNum], nums[dropIndex] = nums[dropIndex], nums[indexOfFirstGreaterDropNum]
	sort.Ints(nums[dropIndex+1:])
}
