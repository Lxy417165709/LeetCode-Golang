package 一维数组

// search 获取旋转排序数组中，目标值的索引。
func search(nums []int, target int) int {
	// 1. 获取旋转排序数组中最小值的索引。
	indexOfMinValue := findIndexOfMinValueFromSortedRotateArray(nums)

	// 2. 拆分为两个排序数组。
	leftSortedArray := nums[:indexOfMinValue]
	rightSortedArray := nums[indexOfMinValue:]

	// 3. 从左数组中，寻找目标值索引。
	indexFoundFromLeftArray := findIndexFromSortedArray(leftSortedArray, target)
	if indexFoundFromLeftArray != -1 {
		return indexFoundFromLeftArray
	}

	// 4. 从右数组中，寻找目标值索引。
	indexFoundFromRightArray := findIndexFromSortedArray(rightSortedArray, target)
	if indexFoundFromRightArray != -1 {
		return indexOfMinValue + indexFoundFromRightArray
	}

	// 5. 未找到。
	return -1
}

// findIndexFromSortedArray 获取排序数组中，目标值的索引。
func findIndexFromSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// findIndexOfMinValueFromSortedRotateArray 获取旋转排序数组中，最小值的索引。 (如果有多个最小值，此时取最左的一个)
func findIndexOfMinValueFromSortedRotateArray(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
